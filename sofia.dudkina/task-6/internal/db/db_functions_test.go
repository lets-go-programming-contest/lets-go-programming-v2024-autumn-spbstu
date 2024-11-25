package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/sssidkn/example_mock/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type MockDatabase struct{}

func (m MockDatabase) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}

type rowTestDb struct {
	names       []string
	uniqueNames []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names:       []string{"Ivan, Gena228"},
		uniqueNames: []string{"Ivan, Gena228"},
	},
	{
		names:       nil,
		uniqueNames: nil,
		errExpected: errors.New("ExpectedError"),
	},
	{
		names:       []string{"Ivan, Gena228, Gena228"},
		uniqueNames: []string{"Ivan, Gena228"},
	},
}

func TestNew(t *testing.T) {
	mockDB := &MockDatabase{}

	service := db.New(mockDB)
	if service.DB != mockDB {
		t.Error("New DB does not match mockDB")
	}
}

func TestGetNames(t *testing.T) {
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
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := db.Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbRows(row.uniqueNames)).
			WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues("name", "users")
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.uniqueNames, names, "row: %d, expected names: %s, actual names: %s", i, row.uniqueNames, names)
	}
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}
