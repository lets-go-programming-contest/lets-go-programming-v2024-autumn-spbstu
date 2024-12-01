package db

import (
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mapset "github.com/deckarep/golang-set"
	"github.com/stretchr/testify/require"
)

type rowTestDb struct {
	names         []string
	namesDistinct []string
	errExpected   error
}

var testTable = []rowTestDb{
	{
		names:         []string{"Nicholas, Adri234, Peter"},
		namesDistinct: []string{"Nicholas, Adri234, Peter"},
	},
	{
		names:         []string{"Nicholas, Adri234, Adri234, Nicholas"},
		namesDistinct: []string{"Nicholas, Adri234"},
	},
	{
		names:         []string{"Nicholas, Peter, Nicholas, Nicholas"},
		namesDistinct: []string{"Nicholas, Peter"},
	},
	{
		names:         nil,
		namesDistinct: nil,
		errExpected:   errors.New("ExpectedError"),
	},
}

func TestNew(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	testDB := New(mockDB)
	dbService := Service{DB: mockDB}
	require.Equal(t, dbService, testDB, "row: %d, expected names: %s, actual names: %s", dbService, testDB)
	require.NotNil(t, dbService, "Expected Service, not nil")
}

func TestGetName(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.GetNames()
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i,
			row.names, names)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockDbRowsDistinct(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues("name", "users")
		if row.errExpected != nil {
			require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %w, actual error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row: %d, names must be nil", i)
			continue
		}
		require.NoError(t, err, "row: %d, error must be nil", i)
		require.Equal(t, row.names, names, "row: %d, expected names: %s, actual names: %s", i, row.namesDistinct, names)
	}
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
	namesSet := mapset.NewSet()
	for _, name := range names {
		namesSet.Add(name)
	}
	for _, name := range namesSet.ToSlice() {
		rows = rows.AddRow(name)
	}
	return rows
}

func TestGetNames_Errors(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	dbService := Service{DB: mockDB}

	tests := []struct {
		name        string
		mockRows    *sqlmock.Rows
		expectedErr string
	}{
		{
			name: "Scan Error",
			mockRows: sqlmock.NewRows([]string{"name"}).
				AddRow(nil),
			expectedErr: "scan rows failedsql: Scan error on column index 0, name \"name\": converting NULL to string is unsupported",
		},
		{
			name: "Rows Error",
			mockRows: sqlmock.NewRows([]string{"name"}).
				AddRow("John").
				RowError(0, fmt.Errorf("row error")),
			expectedErr: "rows error:row error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("SELECT name FROM users").WillReturnRows(tt.mockRows)

			names, err := dbService.GetNames()
			require.Nil(t, names, "Expected names to be nil")
			require.EqualError(t, err, tt.expectedErr, "Expected error did not match")
		})
	}
}

func TestSelectUniqueValues_Errors(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected", err)
	}
	dbService := Service{DB: mockDB}

	tests := []struct {
		name        string
		mockRows    *sqlmock.Rows
		expectedErr string
	}{
		{
			name: "Scan Error",
			mockRows: sqlmock.NewRows([]string{"name"}).
				AddRow(nil),
			expectedErr: "scan rows failedsql: Scan error on column index 0, name \"name\": converting NULL to string is unsupported",
		},
		{
			name: "Rows Error",
			mockRows: sqlmock.NewRows([]string{"name"}).
				AddRow("Alice").
				RowError(0, fmt.Errorf("row error")),
			expectedErr: "rows error:row error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(tt.mockRows)

			values, err := dbService.SelectUniqueValues("name", "users")
			require.Nil(t, values, "Expected values to be nil")
			require.EqualError(t, err, tt.expectedErr, "Expected error did not match")
		})
	}
}
