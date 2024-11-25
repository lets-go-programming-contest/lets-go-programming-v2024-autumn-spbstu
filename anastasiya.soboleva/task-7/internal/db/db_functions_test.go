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

func TestNew(t *testing.T) {
	t.Run("Create Service instance", func(t *testing.T) {
		mockDB, _, err := sqlmock.New()
		require.NoError(t, err)

		service := New(mockDB)
		require.NotNil(t, service)
		require.Equal(t, mockDB, service.DB)
	})
}

func TestGetNamesSuccess(t *testing.T) {
	t.Run("Retrieve multiple names successfully", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		mock.ExpectQuery(selectNamesQuery).
			WillReturnRows(mockDbRows([]string{"Alice", "Bob", "Charlie"}))

		names, err := service.GetNames()

		require.NoError(t, err)
		require.ElementsMatch(t, []string{"Alice", "Bob", "Charlie"}, names)
	})
}

func TestGetNamesQueryError(t *testing.T) {
	t.Run("Query error occurs", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		mock.ExpectQuery(selectNamesQuery).
			WillReturnError(fmt.Errorf("query execution error"))

		names, err := service.GetNames()

		require.Error(t, err)
		require.Nil(t, names)
		require.Contains(t, err.Error(), "query execution error")
	})
}

func TestGetNamesScanError(t *testing.T) {
	t.Run("Scan error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		mock.ExpectQuery(selectNamesQuery).
			WillReturnRows(mockDbRowsWithScanError())

		names, err := service.GetNames()

		require.Error(t, err)
		require.Nil(t, names)
		require.Contains(t, err.Error(), "failed to scan row")
	})
}

func TestGetNamesIterationError(t *testing.T) {
	t.Run("Iteration error", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		mock.ExpectQuery(selectNamesQuery).
			WillReturnRows(mockDbRowsWithIterationError())

		names, err := service.GetNames()

		require.Error(t, err)
		require.Nil(t, names)
		require.Contains(t, err.Error(), "row iteration error")
	})
}

func TestSelectUniqueValuesSuccess(t *testing.T) {
	t.Run("Select unique values", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		query := fmt.Sprintf(selectUniqueValuesQuery, "name", "users")
		mock.ExpectQuery(query).
			WillReturnRows(mockDbRows([]string{"Bob", "Charlie"}))

		uniqueNames, err := service.SelectUniqueValues("name", "users")

		require.NoError(t, err)
		require.ElementsMatch(t, []string{"Bob", "Charlie"}, uniqueNames)
	})
}

func TestSelectUniqueValuesQueryError(t *testing.T) {
	t.Run("Query error on unique values", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		query := fmt.Sprintf(selectUniqueValuesQuery, "name", "users")
		mock.ExpectQuery(query).
			WillReturnError(fmt.Errorf("query execution error"))

		uniqueNames, err := service.SelectUniqueValues("name", "users")

		require.Error(t, err)
		require.Nil(t, uniqueNames)
		require.Contains(t, err.Error(), "query execution error")
	})
}

func TestSelectUniqueValuesScanError(t *testing.T) {
	t.Run("Scan error on unique values", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		query := fmt.Sprintf(selectUniqueValuesQuery, "name", "users")
		mock.ExpectQuery(query).
			WillReturnRows(mockDbRowsWithScanError())

		uniqueNames, err := service.SelectUniqueValues("name", "users")

		require.Error(t, err)
		require.Nil(t, uniqueNames)
		require.Contains(t, err.Error(), "failed to scan row")
	})
}

func TestSelectUniqueValuesIterationError(t *testing.T) {
	t.Run("Iteration error on unique values", func(t *testing.T) {
		mockDB, mock, err := sqlmock.New()
		require.NoError(t, err)

		service := Service{DB: mockDB}
		query := fmt.Sprintf(selectUniqueValuesQuery, "name", "users")
		mock.ExpectQuery(query).
			WillReturnRows(mockDbRowsWithIterationError())

		uniqueNames, err := service.SelectUniqueValues("name", "users")

		require.Error(t, err)
		require.Nil(t, uniqueNames)
		require.Contains(t, err.Error(), "row iteration error")
	})
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
