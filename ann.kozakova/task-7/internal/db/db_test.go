package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names: []string{"Ivan, Gena228"},
	},
	{
		names:       nil,
		errExpected: errors.New("ExpectedError"),
	},
	{
		names: []string{"Ivan, Gena228", "Ivan"},
	},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := DBService{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).
			WillReturnError(row.errExpected)
		names, err := dbService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error:%w, actual error: %w", i,
				row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i,
			row.names, names)
	}
}
func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestSetUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := DBService{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbUniqueRows(row.names)).
			WillReturnError(row.errExpected)
		uniqueNames, err := dbService.SelectUniqueValues("name", "users")
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error:%w, actual error: %w", i,
				row.errExpected, err)
			require.Nil(t, uniqueNames, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, uniqueNames, "row: %d, expected names: %s, actual names: %s", i,
			row.names, uniqueNames)
	}
}

func mockDbUniqueRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	uniqueSet := mapset.NewSet()
	for _, str := range names {
		uniqueSet.Add(str)
	}
	for _, name := range uniqueSet.ToSlice() {
		rows = rows.AddRow(name)
	}
	return rows
}
