package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names         []string
	namesDistinct []string
	errExpected   error
}

var testTable = []rowTestDb{
	{
		names:         nil,
		namesDistinct: nil,
		errExpected:   errors.New("ExpectedError"),
	},
	{
		names:         []string{"Alina, Gena, Elena"},
		namesDistinct: []string{"Alina, Gena, Elena"},
	},
	{
		names:         []string{"Alina, Gena, Nata, Elena"},
		namesDistinct: []string{"Alina, Elena"},
	},
	{
		names:         []string{"Alina, Gena"},
		namesDistinct: []string{"Alina, Gena, Elena"},
	},
}

func TestNew(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	testDB := New(mockDB)
	dbService := Service{DB: mockDB}
	require.Equal(t, dbService, testDB, "row: %d, expected names: %s, actual names: %s", dbService, testDB)
	require.NotNil(t, dbService, "Expected Service, not nil")
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i,
			row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbRowsDistinct(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues("name", "users")
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.namesDistinct, names)
	}
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func mockDbRowsDistinct(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	namesSet := mapset.NewSet()
	for _, name := range names {
		namesSet.Add(name)
	}
	for _, name := range namesSet.ToSlice() {
		rows = rows.AddRow(name)
	}
	return rows
}

func mockDbRowsWithScanError(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}

	rows.AddRow(nil)
	return rows
}

func TestGetNameWithScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := Service{DB: mockDB}

	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRowsWithScanError([]string{"Alina", "Gena", "Elena"}))
	names, err := dbService.GetNames()

	require.Error(t, err, "Expected scan error")
	require.Nil(t, names, "Names should be nil if scan fails")
}

func TestSelectUniqueValuesWithScanError(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := Service{DB: mockDB}

	mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbRowsWithScanError([]string{"Alina", "Gena", "Elena"}))
	names, err := dbService.SelectUniqueValues("name", "users")

	require.Error(t, err, "Expected scan error")
	require.Nil(t, names, "Names should be nil if scan fails")
}
