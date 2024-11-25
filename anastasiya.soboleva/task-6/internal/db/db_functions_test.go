package db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	selectNamesQuery        = "SELECT name FROM users"
	selectUniqueValuesQuery = "SELECT DISTINCT %s FROM %s"
)

type rowTestDb struct {
	names         []string
	namesDistinct []string
	errExpected   error
}

var testTable = []rowTestDb{
	{
		names:         []string{"Alice", "Bob", "Charlie", "Charlie"},
		namesDistinct: []string{"Alice", "Bob", "Charlie"},
		errExpected:   nil,
	},
	{
		names:         nil,
		namesDistinct: nil,
		errExpected:   fmt.Errorf("query execution error"),
	},
	{
		names:         nil,
		namesDistinct: nil,
		errExpected:   fmt.Errorf("failed to scan row"),
	},
	{
		names:         nil,
		namesDistinct: nil,
		errExpected:   fmt.Errorf("row iteration error"),
	},
}

func TestNew(t *testing.T) {
	t.Run("Create Service instance", func(t *testing.T) {
		mockDB, _, err := sqlmock.New()
		require.NoError(t, err)

		service := New(mockDB)
		require.NotNil(t, service)
		require.Equal(t, mockDB, service.DB)
	})
}

func TestGetNames(t *testing.T) {
	t.Parallel()
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test for names: %v", testCase.names), func(t *testing.T) {
			t.Parallel()
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			service := Service{DB: mockDB}

			if testCase.errExpected == nil {
				mock.ExpectQuery(selectNamesQuery).
					WillReturnRows(mockDbRows(testCase.names))

				names, err := service.GetNames()

				require.NoError(t, err)
				require.ElementsMatch(t, testCase.names, names)
			} else if testCase.errExpected.Error() == "query execution error" {
				mock.ExpectQuery(selectNamesQuery).
					WillReturnError(testCase.errExpected)

				names, err := service.GetNames()

				require.Error(t, err)
				require.Nil(t, names)
				require.Contains(t, err.Error(), "query execution error")
			} else if testCase.errExpected.Error() == "failed to scan row" {
				mock.ExpectQuery(selectNamesQuery).
					WillReturnRows(mockDbRowsWithScanError())

				names, err := service.GetNames()

				require.Error(t, err)
				require.Nil(t, names)
				require.Contains(t, err.Error(), "failed to scan row")
			} else if testCase.errExpected.Error() == "row iteration error" {
				mock.ExpectQuery(selectNamesQuery).
					WillReturnRows(mockDbRowsWithIterationError())

				names, err := service.GetNames()

				require.Error(t, err)
				require.Nil(t, names)
				require.Contains(t, err.Error(), "row iteration error")
			}
		})
	}
}

func TestSelectUniqueValues(t *testing.T) {
	t.Parallel()
	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("Test for distinct values: %v", testCase.namesDistinct), func(t *testing.T) {
			t.Parallel()
			mockDB, mock, err := sqlmock.New()
			require.NoError(t, err)

			service := Service{DB: mockDB}
			query := fmt.Sprintf(selectUniqueValuesQuery, "name", "users")

			if testCase.errExpected == nil {
				mock.ExpectQuery(query).
					WillReturnRows(mockDbRows(testCase.namesDistinct))

				uniqueNames, err := service.SelectUniqueValues("name", "users")

				require.NoError(t, err)
				require.ElementsMatch(t, testCase.namesDistinct, uniqueNames)
			} else if testCase.errExpected.Error() == "query execution error" {
				mock.ExpectQuery(query).
					WillReturnError(testCase.errExpected)

				uniqueNames, err := service.SelectUniqueValues("name", "users")

				require.Error(t, err)
				require.Nil(t, uniqueNames)
				require.Contains(t, err.Error(), "query execution error")
			} else if testCase.errExpected.Error() == "failed to scan row" {
				mock.ExpectQuery(query).
					WillReturnRows(mockDbRowsWithScanError())

				uniqueNames, err := service.SelectUniqueValues("name", "users")

				require.Error(t, err)
				require.Nil(t, uniqueNames)
				require.Contains(t, err.Error(), "failed to scan row")
			} else if testCase.errExpected.Error() == "row iteration error" {
				mock.ExpectQuery(query).
					WillReturnRows(mockDbRowsWithIterationError())

				uniqueNames, err := service.SelectUniqueValues("name", "users")

				require.Error(t, err)
				require.Nil(t, uniqueNames)
				require.Contains(t, err.Error(), "row iteration error")
			}
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

func mockDbRowsWithScanError() *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	rows.AddRow(nil)
	return rows
}

func mockDbRowsWithIterationError() *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"}).
		AddRow("Alice").
		AddRow("Bob")
	rows.RowError(1, fmt.Errorf("row iteration error"))
	return rows
}
