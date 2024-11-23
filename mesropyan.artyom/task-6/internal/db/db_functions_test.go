package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	namesUnique []string

	errExpected error
}

var testTable = []rowTestDb{
	{
		names:       []string{"Ivan", "Gena228", "Andrey"},
		namesUnique: []string{"Ivan", "Gena228", "Andrey"},
	},
	{
		names:       []string{"Ivan", "Gena228", "Gena228"},
		namesUnique: []string{"Ivan", "Gena228"},
	},
	{
		names:       nil,
		namesUnique: nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func TestNew(t *testing.T) {

	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	testDB := New(mockDB)
	Service := Service{DB: mockDB}
	require.Equal(t, Service, testDB,
		"row: %d, expected names: %s, actual names: %s", Service, testDB)
	require.NotNil(t, Service, "Should be Service, not nil")
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	Service := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).
			WillReturnError(row.errExpected)
		names, err := Service.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected,
				"row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names,
			"row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	Service := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT (.+) FROM (.+)").WillReturnRows(mockDbRowsUnique(row.names)).
			WillReturnError(row.errExpected)
		names, err := Service.SelectUniqueValues("names", "users")
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected,
				"row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.namesUnique, names,
			"row: %d, expected names: %s, actual names: %s", i, row.namesUnique, names)
	}
}
func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func mockDbRowsUnique(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	uniqueNames := make(map[string]bool)

	for _, name := range names {
		if !uniqueNames[name] {
			uniqueNames[name] = true
			rows = rows.AddRow(name)
		}
	}
	return rows
}
