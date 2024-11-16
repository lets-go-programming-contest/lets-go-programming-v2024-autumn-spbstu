package db

import (
	"database/sql"
	"fmt"
)

type Database interface {
	Query(query string, args ...any) (*sql.Rows, error)
}

type Service struct {
	DB Database
}

func New(db Database) Service {
	return Service{DB: db}
}

func (service Service) GetNames() ([]string, error) {
	query := "SELECT name FROM users"

	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query DB err: %w", err)
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err = rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("rows.Scan err: %w", err)
		}
		names = append(names, name)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err err: %w", err)
	}

	return names, nil
}

func (service Service) SelectUniqueValues(columnName string, tableName string) ([]string, error) {
	query := "SELECT DISTINCT " + columnName + " FROM " + tableName
	var err error
	rows, err := service.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query DB err: %w", err)
	}
	defer rows.Close()

	values := make([]string, 0)
	for rows.Next() {
		var value string
		if err = rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("rows.Scan err: %w", err)
		}
		values = append(values, value)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows.Err err: %w", err)
	}

	return values, nil
}
