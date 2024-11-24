package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/Mmmakskl/task-7/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type testTableDB struct {
	Names         []string
	ExpectedValue []string
	ExpectedError error
}

var testTable = []testTableDB{
	{
		Names:         []string{"Maksim", "Ivan", "Andrew"},
		ExpectedValue: []string{"Maksim", "Ivan", "Andrew"},
		ExpectedError: nil,
	},
	{
		Names:         []string{"Maksim", "Ivan", "Andrew", "Ivan"},
		ExpectedValue: []string{"Maksim", "Ivan", "Andrew"},
		ExpectedError: nil,
	},
	{
		Names:         nil,
		ExpectedValue: nil,
		ExpectedError: errors.New("ExpectedError"),
	},
	{
		Names:         []string{},
		ExpectedValue: []string{},
		ExpectedError: nil,
	},
}

func TestGetNames(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshalling expected json data", err)
	}

	dbService := db.Service{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.Names)).WillReturnError(row.ExpectedError)

		names, err := dbService.GetNames()

		if row.ExpectedError != nil {
			require.ErrorIs(t, err, row.ExpectedError, "row: %d, expected error:%w, actual error: %w", i, row.ExpectedError, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.Names, emptyIfNil(names), "row: %d, expected names: %s, actual names: %s", i, row.Names, names)
	}

	t.Run("rows.Scan() handling", func(t *testing.T) {
		testScan(t, dbService, mock, "SELECT name FROM users")
	})

	t.Run("rows.Err() handling", func(t *testing.T) {
		testRowsErr(t, dbService, mock, "SELECT name FROM users")
	})
}

func TestSelectUniqueValues(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshalling expected json data", err)
	}
	defer mockDB.Close()

	dbService := db.Service{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbRowsDistinct(row.Names)).WillReturnError(row.ExpectedError)

		names, err := dbService.SelectUniqueValues("name", "users")

		if row.ExpectedError != nil {
			require.ErrorIs(t, err, row.ExpectedError, "row: %d, expected error:%w, actual error: %w", i, row.ExpectedError, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.ElementsMatch(t, row.ExpectedValue, names, "expected uniaue values: %v, actual values: %v", row.ExpectedValue, names)
	}

	t.Run("rows.Scan() handling", func(t *testing.T) {
		testScan(t, dbService, mock, "SELECT DISTINCT name FROM users")
	})

	t.Run("rows.Err() handling", func(t *testing.T) {
		testRowsErr(t, dbService, mock, "SELECT DISTINCT name FROM users")
	})
}

func testRowsErr(t *testing.T, dbService db.Service, mock sqlmock.Sqlmock, query string) {
	t.Helper()

	var (
		names []string
		err   error
	)

	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Maksim").RowError(0, errors.New("row processing error")))

	if query == "SELECT name FROM users" {
		names, err = dbService.GetNames()
	} else {
		names, err = dbService.SelectUniqueValues("name", "users")
	}

	require.Error(t, err, "rows.Err() should return an error")
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows.Err() failed", "unexpected error message for rows.Err()")
}

func testScan(t *testing.T, dbService db.Service, mock sqlmock.Sqlmock, query string) {
	t.Helper()

	var (
		names []string
		err   error
	)

	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(nil))

	if query == "SELECT name FROM users" {
		names, err = dbService.GetNames()
	} else {
		names, err = dbService.SelectUniqueValues("name", "users")
	}

	require.Error(t, err, "rows.Scan() should return an error")
	require.Nil(t, names)
	require.Contains(t, err.Error(), "rows.Scan() failed", "unexpected error message for rows.Scan()")
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB := &mockDbStruct{}

	service := db.New(mockDB)

	require.NotNil(t, service)
	require.Equal(t, mockDB, service.DB, "Service should have the provided DB")
}

type mockDbStruct struct{}

func (db *mockDbStruct) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}

func mockDbRowsDistinct(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	uniqueNames := make(map[string]bool)

	for _, name := range names {
		if !uniqueNames[name] {
			uniqueNames[name] = true

			rows.AddRow(name)
		}
	}
	return rows
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func emptyIfNil(names []string) []string {
	if names == nil {
		return []string{}
	}
	return names
}
