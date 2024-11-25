package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kirill.romanchuk/task-7/internal/db"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names       []string
	uniqNames   []string
	errExpected error
}

var testTable = []rowTestDb{
	// Тест с несколькими уникальными именами
	{
		names:       []string{"Ivan", "Gena228", "Anna", "Anna", "lox"},
		uniqNames:   []string{"Ivan", "Gena228", "Anna", "lox"},
		errExpected: nil,
	},
	// Тест, где возвращается только одно имя
	{
		names:       []string{"Solo"},
		uniqNames:   []string{"Solo"},
		errExpected: nil,
	},
	// Тест с пустым массивом имен
	{
		names:       nil,
		uniqNames:   nil,
		errExpected: nil,
	},
	// Тест с дублирующимися именами (все имена одинаковые)
	{
		names:       []string{"Duplicate", "Duplicate", "Duplicate"},
		uniqNames:   []string{"Duplicate"},
		errExpected: nil,
	},
	// Тест для обработки ошибки базы данных
	{
		names:       []string{"Ivan", "Gena228"},
		uniqNames:   nil,
		errExpected: errors.New("ExpectedError"),
	},
	// Тест с многоразовыми именами
	{
		names:       []string{"Alice", "Bob", "Alice", "Bob", "Charlie"},
		uniqNames:   []string{"Alice", "Bob", "Charlie"}, // Уникальные имена
		errExpected: nil,
	},
	// Тест с большими именами (включая пробелы)
	{
		names:       []string{"John Doe", "Jane Doe", "John"},
		uniqNames:   []string{"John Doe", "Jane Doe"},
		errExpected: nil,
	},
	// Тест с пустым массивом имен и ошибкой
	{
		names:       nil,
		uniqNames:   nil,
		errExpected: errors.New("ExpectedError"),
	},
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := db.New(mockDB)

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
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	dbService := db.New(mockDB)
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").
			WillReturnRows(mockDbRows(row.uniqNames)).
			WillReturnError(row.errExpected)

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
