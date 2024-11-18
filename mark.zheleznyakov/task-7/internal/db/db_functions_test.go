package db_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	myDb "github.com/mrqiz/task-6/internal/db"
)

func TestGetNames(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("got an '%s' when opening a stub database connection", err)
	}
	defer db.Close()

	service := myDb.New(db)

	rows := sqlmock.NewRows([]string{"name"}).AddRow("Nikolay").AddRow("Mark")
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(rows)

	names, err := service.GetNames()
	assert.NoError(t, err)
	assert.Equal(t, []string{"Nikolay", "Mark"}, names)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("something is meh: %s", err)
	}
}

func TestSelectUniqueValues(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' occurred when opening a stub database connection", err)
	}
	defer db.Close()

	service := myDb.New(db)

	rows := sqlmock.NewRows([]string{"columnName"}).AddRow("Foo").AddRow("Bar")
	mock.ExpectQuery("SELECT DISTINCT columnName FROM tableName").WillReturnRows(rows)

	values, err := service.SelectUniqueValues("columnName", "tableName")
	assert.NoError(t, err)
	assert.Equal(t, []string{"Foo", "Bar"}, values)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("it got worse: %s", err)
	}
}

func TestGetNamesError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("'%s' when opening a stub database connection", err)
	}
	defer db.Close()

	service := myDb.New(db)

	mock.ExpectQuery("SELECT name FROM users").WillReturnError(errors.New("database error"))

	names, err := service.GetNames()
	assert.Error(t, err)
	assert.Nil(t, names)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("it got worse: %s", err)
	}
}

func TestSelectUniqueValuesError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("'%s' when opening a stub database connection", err)
	}
	defer db.Close()

	service := myDb.New(db)

	mock.ExpectQuery("SELECT DISTINCT columnName FROM tableName").WillReturnError(errors.New("database error"))

	values, err := service.SelectUniqueValues("columnName", "tableName")
	assert.Error(t, err)
	assert.Nil(t, values)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("it got worse: %s", err)
	}
}
