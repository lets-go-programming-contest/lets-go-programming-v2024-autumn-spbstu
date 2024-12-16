package generate

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	mapset "github.com/deckarep/golang-set/v2"
)

type TestDB struct {
	names       []string
	errExpected error
}

var testTable = []TestDB{
	{
		names:       []string{"Masha", "Danil", "Pasha"},
		errExpected: nil,
	},
	{
		names:       []string{"Danil", "Dima", "Dima"},
		errExpected: nil,
	},
	{
		names:       nil,
		errExpected: errors.New("no names"),
	},
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
		if err != row.errExpected {
			t.Errorf("Test %d: Expected error '%s', got '%s'", i, row.errExpected, err)
		}
		if len(names) != len(row.names) {
			t.Errorf("Test %d: Expected %d names, got %d", i, len(row.names), len(names))
		}
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
		if err != row.errExpected {
			t.Errorf("Test %d: Expected error '%s', got '%s'", i, row.errExpected, err)
		}
		if len(names) != len(row.names) {
			t.Errorf("Test %d: Expected %d names, got %d", i, len(row.names), len(names))
		}
	}
}
