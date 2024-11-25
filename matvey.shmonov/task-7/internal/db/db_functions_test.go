package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Koshsky/task-6/internal/db"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	uniqNames   []string
	errExpected error
}

func newRowsFromStrings(column string, values []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{column})
	for _, value := range values {
		rows.AddRow(value)
	}
	return rows
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
	{
		names:       []string{"Alice", "Bob", "Alice", "Charlie"},
		uniqNames:   []string{"Alice", "Bob", "Charlie"},
		errExpected: nil,
	},
	{
		names:       []string{"John", "Doe", "John", "Doe"},
		uniqNames:   []string{"John", "Doe"},
		errExpected: nil,
	},
	{
		names:       []string{"Zara", "Mike", "Zara", "Anna", "Mike"},
		uniqNames:   []string{"Zara", "Mike", "Anna"},
		errExpected: nil,
	},
	{
		names:       []string{"One", "Two", "Three", "One", "Two"},
		uniqNames:   []string{"One", "Two", "Three"},
		errExpected: nil,
	},
	{
		names:       []string{"A", "B", "C", "A", "B", "C"},
		uniqNames:   []string{"A", "B", "C"},
		errExpected: nil,
	},
	{
		names:       []string{"Test", "", "Test"},
		uniqNames:   []string{"Test"},
		errExpected: nil,
	},
	{
		names:       []string{"Unique1", "Unique2", "Unique3"},
		uniqNames:   []string{"Unique1", "Unique2", "Unique3"},
		errExpected: nil,
	},
	{
		names:       []string{"Same", "Same", "Same", "Same"},
		uniqNames:   []string{"Same"},
		errExpected: nil,
	},
	{
		names:       []string{"A", "B", "C", "D", "E"},
		uniqNames:   []string{"A", "B", "C", "D", "E"},
		errExpected: nil,
	},
	{
		names:       []string{"Duplicate", "Duplicate", "Unique"},
		uniqNames:   []string{"Duplicate", "Unique"},
		errExpected: nil,
	},
}

func TestGetNames(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer mockDb.Close()

	service := db.New(mockDb)

	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(newRowsFromStrings("name", row.names)).WillReturnError(row.errExpected)
		names, err := service.GetNames()

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
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a mock database connection", err)
	}
	defer mockDb.Close()

	service := db.New(mockDb)

	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(newRowsFromStrings("name", row.uniqNames)).WillReturnError(row.errExpected)
		names, err := service.SelectUniqueValues("name", "users")

		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}

		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.uniqNames, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}
