package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"

	"github.com/Piyavva/task-6/internal/db"
)

type rowDb struct {
	names       []string
	errExpected error
}

var testTable = []rowDb{
	{
		[]string{"Anton", "Danil"}, nil,
	},
	{
		nil, errors.New("test error"),
	},
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := db.Service{DB: mockDB}
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

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := db.Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT column FROM table").WillReturnRows(mockURows(row.names)).
			WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues("column", "table")
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

func mockURows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	mp := make(map[string]int)
	for _, name := range names {
		if _, ok := mp[name]; !ok {
			rows = rows.AddRow(name)
			mp[name] = mp[name] + 1
		}
	}
	return rows
}
