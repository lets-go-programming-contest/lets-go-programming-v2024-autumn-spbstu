package db_test

import (
	"errors"
	"testing"

	"task-7/internal/db"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

// go test -coverpkg=./... -coverprofile cover.out ./...
// go tool cover -html cover.out -o cover.html

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
}

func TestGetName(t *testing.T) {
	t.Parallel()

	t.Run("success GetNames()", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
		}
		dbService := db.Service{DB: mockDB}

		for i, row := range testTable {
			mock.ExpectQuery("SELECT name FROM users").
				WillReturnRows(mockDbRows(row.names)).
				WillReturnError(row.errExpected)
			names, err := dbService.GetNames()
			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w ",
					i,
					row.errExpected,
					err,
				)
				require.Nil(t, names, "row: %d, names must be nil", i)
				continue
			}
			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s ", i,
				row.names, names)
		}
	})

	t.Run("rows.Scan(&name) err", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
		}
		dbService := db.New(mockDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow(nil))
		_, err = dbService.GetNames()

		require.Error(t, err, "error must be not nil, got %v", err)
	})

	t.Run("rows.Err()", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
		}
		dbService := db.New(mockDB)

		mock.ExpectQuery("SELECT name FROM users").
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("one").
				AddRow("two").
				RowError(1, sqlmock.ErrCancelled),
			)

		_, err = dbService.GetNames()

		require.Error(t, err, "error must be not nil, got %v", err)
	})
}

func TestSelectUniqueValues(t *testing.T) {
	t.Parallel()

	t.Run("success SelectUniqueValues()", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
		}
		dbService := db.Service{DB: mockDB}

		for i, row := range testTable {
			mock.ExpectQuery("SELECT DISTINCT name FROM users").
				WillReturnRows(mockDbRows(row.names)).
				WillReturnError(row.errExpected)
			names, err := dbService.SelectUniqueValues("name", "users")
			if row.errExpected != nil {
				require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w ",
					i,
					row.errExpected,
					err,
				)
				require.Nil(t, names, "row: %d, names must be nil", i)
				continue
			}
			require.NoError(t, err, "row: %d, error must be nil", i)
			require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s ", i,
				row.names, names)
		}
	})

	t.Run("rows.Scan(&value) err", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
		}
		dbService := db.New(mockDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow(nil))

		_, err = dbService.SelectUniqueValues("name", "users")

		require.Error(t, err, "error must be not nil, got %v", err)
	})

	t.Run("rows.Err()", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
		}
		dbService := db.New(mockDB)

		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(sqlmock.NewRows([]string{"name"}).
				AddRow("one").
				AddRow("two").
				RowError(1, sqlmock.ErrCancelled),
			)

		_, err = dbService.SelectUniqueValues("name", "users")

		require.Error(t, err, "error must be not nil, got %v", err)
	})
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}
