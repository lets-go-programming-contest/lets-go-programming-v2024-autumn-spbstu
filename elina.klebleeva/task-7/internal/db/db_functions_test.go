package db_test

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/EmptyInsid/task-7/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	uniqeExpect []string
	errExpected error
}

var testTable = []rowTestDb{
	{
		names:       []string{"Ivan", "Gens228"},
		uniqeExpect: []string{"Ivan", "Gens228"},
		errExpected: nil,
	},
	{
		names:       nil,
		uniqeExpect: nil,
		errExpected: errors.New("ExpectedError"),
	},
	{
		names:       []string{"Ivan", "Gens228", "Ivan"},
		uniqeExpect: []string{"Ivan", "Gens228"},
		errExpected: nil,
	},
	{
		names:       []string{},
		uniqeExpect: []string{},
		errExpected: nil,
	},
}

type MockDatabase struct{}

func (m *MockDatabase) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func TestNew(t *testing.T) {
	t.Parallel()

	mockDB := &MockDatabase{}

	service := db.New(mockDB)

	require.NotNil(t, service)
	require.Equal(t, mockDB, service.DB, "DB field in DBService should be set to the mockDB instance")
}

func TestGetName(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshalling expected json data", err)
	}
	dbService := db.Service{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(mockDbRows(row.names)).
			WillReturnError(row.errExpected)

		names, err := dbService.GetNames()

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, ensureNotNil(names), "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening mock database connection: %v", err)
	}
	defer mockDB.Close()

	dbService := db.Service{DB: mockDB}

	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(mockDbRowsDistinct(row.names)).
			WillReturnError(row.errExpected)

		names, err := dbService.SelectUniqueValues("name", "users")

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %v, actual error: %v", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names should be nil when an error is expected", i)
			continue
		}

		require.NoError(t, err, "row: %d, error should be nil", i)
		require.ElementsMatch(t, row.uniqeExpect, names, "expected unique values: %v, got: %v", row.uniqeExpect, names)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %v", err)
	}
}

func TestScanError(t *testing.T) {
	t.Parallel()

	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("unexpected error when opening mock database connection: %v", err)
	}
	defer mockDB.Close()

	dbService := db.Service{DB: mockDB}

	// ---------- Test SelectUniqueValues ----------

	queryUnique := "SELECT DISTINCT name FROM users"
	rows := sqlmock.NewRows([]string{"name"}).AddRow(nil) // `nil` create error for scan

	mock.ExpectQuery(queryUnique).WillReturnRows(rows)

	values, err := dbService.SelectUniqueValues("name", "users")
	require.Error(t, err, "expected an error from rows.Scan in SelectUniqueValues")
	require.Contains(t, err.Error(), "rows.Scan failed", "unexpected error message for SelectUniqueValues")
	require.Nil(t, values, "expected values to be nil due to scan error in SelectUniqueValues")

	// ---------- Test GetName ----------

	queryGetName := "SELECT name FROM users"
	mock.ExpectQuery(queryGetName).WithArgs().WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(nil))

	name, err := dbService.GetNames()
	require.Error(t, err, "expected an error from row.Scan in GetName")
	require.Contains(t, err.Error(), "rows.Scan failed", "unexpected error message for GetName")
	require.Equal(t, []string([]string(nil)), name, "expected name to be empty due to scan error in GetName")
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
	uniqueNames := make(map[string]bool)

	for _, name := range names {
		if !uniqueNames[name] {
			uniqueNames[name] = true
			rows = rows.AddRow(name)
		}
	}
	return rows
}

func ensureNotNil(s []string) []string {
	if s == nil {
		return []string{}
	}
	return s
}
