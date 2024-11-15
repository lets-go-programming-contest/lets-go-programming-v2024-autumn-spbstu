package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/require"
)

type TestDB struct {
	names       []string
	errExpected error
}

var testTable = []TestDB{
	{
		nil, nil, //[]string{"Masha", "Danil", "Pasha"}
	},
	{
		nil, nil, //errors.New("test error")
	},
	/*{
		names:       nil,
		errExpected: errors.New("no names"),
	},*/
}

func mockDbRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	for _, name := range names {
		rows.AddRow(name)
	}
	return rows
}

func TestDBServiceGetNames(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT name FROM users").WillReturnRows(mockDbRows(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.GetNames()
		if row.errExpected == nil {
			require.ErrorIs(t, err, row.errExpected, "row %d: expected error: %w, current error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row %d: expected names to be nil", i)
			continue
		}
		require.NoError(t, err, "row %d: expected error to be nil", i)
		require.Equal(t, row.names, names, "row %d: expected names to be equal", i)
	}
}

func mockUniqueRows(names []string) *sqlmock.Rows {
	rows := sqlmock.NewRows([]string{"name"})
	mySet := mapset.NewSet[string]()
	for _, name := range names {
		mySet.Add(name)
	}
	for _, name := range mySet.ToSlice() {
		rows.AddRow(name)
	}
	return rows
}

func TestSelectUniqueValues(t *testing.T) {
	t.Parallel()
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dbService := Service{DB: mockDB}
	for i, row := range testTable {
		mock.ExpectQuery("SELECT DISTINCT name FROM users").WillReturnRows(mockUniqueRows(row.names)).WillReturnError(row.errExpected)
		names, err := dbService.SelectUniqueValues("name", "users")
		if row.errExpected == nil {
			require.ErrorIs(t, err, row.errExpected, "row %d: expected error: %w, current error: %w", i, row.errExpected, err)
			require.Nil(t, names, "row %d: expected names to be nil", i)
			continue
		}
		require.NoError(t, err, "row %d: expected error to be nil", i)
		require.Equal(t, row.names, names, "row %d: expected names to be equal", i)
	}
}

// go test -coverprofile=profile
// go tool cover -html=profile
