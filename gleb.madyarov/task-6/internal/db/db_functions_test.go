package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/Madyarov-Gleb/task-6/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	uniqNames   []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names:       []string{"Ivan", "Gena", "Ivan"},
		uniqNames:   []string{"Ivan", "Gena"},
		errExpected: nil,
	},
	{
		names:       nil,
		uniqNames:   nil,
		errExpected: errors.New("Error"),
	},
	{
		names:       []string{"Ivan", "Gena", "Nikita"},
		uniqNames:   []string{"Ivan", "Gena", "Nikita"},
		errExpected: nil,
	},
}

type MockDatabase struct{}

func TestGetName(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := db.Service{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
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
	t.Parallel()

	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error sqlmock.New: %w", err)
	}
	dbService := db.Service{DB: mockDb}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbRows(row.uniqNames)).WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues("name", "users")

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.uniqNames, names, "row: %d, expected names: %s, actual names: %s", i, row.uniqNames, names)
	}
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}

	return rows
}

func (m *MockDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDb := &MockDatabase{}
	service := db.New(mockDb)
	require.NotNil(t, service)
	require.Equal(t, mockDb, service.DB, "Db must be equal mockDb")

}
