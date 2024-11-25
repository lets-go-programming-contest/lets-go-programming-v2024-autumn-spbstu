package db_test

import (
	"errors"
	"fmt"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"

	"github.com/stretchr/testify/require"

	"erdem.istaev/task-7/internal/db"
)

type RowTestDB struct {
	names       []string
	uniqueNames []string
	errExpected error
}

var testTable = []RowTestDB{
	{
		names:       []string{"Ivan", "Gena228"},
		uniqueNames: []string{"Ivan", "Gena228"},
		errExpected: nil,
	},
	{
		names:       []string{"Ivan", "Gena228", "Ivan"},
		uniqueNames: []string{"Ivan", "Gena228"},
		errExpected: nil,
	},
	{
		names:       nil,
		errExpected: errors.New("No rows in result set"),
	},
}

func TestGetNames(t *testing.T) {
	for i, row := range testTable {
		t.Run(fmt.Sprintf("TestCase%d", i), func(t *testing.T) {
			t.Parallel()

			mockDb, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer mockDb.Close()

			dbService := db.Service{DB: mockDb}

			mock.ExpectQuery("SELECT name FROM users").
				WillReturnRows(mockDbRows(row.names)).
				WillReturnError(row.errExpected)

			names, err := dbService.GetNames()
			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected,
					"row: %d, expected error:%w, actual error: %w",
					i, row.errExpected, err)
				require.Nil(t, names, "row: %d, names must be nil", i)

				return
			}

			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.names, names,
				"row: %d, expected names: %s, actual names: %s",
				i, row.names, names)
		})
	}
}

func TestSelectUniqueValues(t *testing.T) {
	for i, row := range testTable {
		t.Run(fmt.Sprintf("TestCase%d", i), func(t *testing.T) {
			t.Parallel()

			mockDb, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer mockDb.Close()

			dbService := db.Service{DB: mockDb}

			mock.ExpectQuery("SELECT DISTINCT name FROM users").
				WillReturnRows(mockDbUniqRows(row.names)).
				WillReturnError(row.errExpected)

			names, err := dbService.SelectUniqueValues("name", "users")
			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected,
					"row: %d, expected error:%w, actual error: %w",
					i, row.errExpected, err)
				require.Nil(t, names, "row: %d, names must be nil", i)
				return
			}

			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.uniqueNames, names,
				"row: %d, expected names: %s, actual names: %s",
				i, row.uniqueNames, names)
		})
	}
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}

	return rows
}

func mockDbUniqRows(names []string) *sqlmock.Rows {
	uniqueNames := make(map[string]bool)
	for _, name := range names {
		uniqueNames[name] = true
	}

	rows := sqlmock.NewRows([]string{"name"})
	for name := range uniqueNames {
		rows.AddRow(name)
	}

	return rows
}
