package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"task-9/internal/contact"
)

type PgSQLRepository struct {
	DB *sql.DB
}

func NewPgSQLController(uName, uPass, host, dbName string, port int) (PgSQLRepository, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		uName, uPass, host, port, dbName)

	repo := &PgSQLRepository{}
	var err error

	repo.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return *repo, fmt.Errorf("%w + pgsql; %w", ErrOpenDB, err)
	}

	err = repo.Ping()
	if err != nil {
		err_ := repo.Close()
		if err_ != nil {
			return PgSQLRepository{}, errors.Join(err, err_)
		}
		return *repo, fmt.Errorf("%w + pgsql; %w", ErrPingDB, err)
	}

	err = repo.PrepareDBContent()
	if err != nil {
		err_ := repo.Close()
		if err_ != nil {
			return PgSQLRepository{}, errors.Join(err, err_)
		}
		return *repo, err
	}

	return *repo, nil
}

func (m PgSQLRepository) Ping() error {
	return m.DB.Ping()
}

func (m PgSQLRepository) Close() error {
	return m.DB.Close()
}

func (m PgSQLRepository) PrepareDBContent() error {
	qs := []string{
		`DROP TABLE IF EXISTS contacts;`,

		`CREATE TABLE contacts (
    		ID SERIAL PRIMARY KEY,
    		Name VARCHAR(255) PRIMARY KEY,
    		Phone VARCHAR(255) NOT NULL,
		);`,
	}

	for _, q := range qs {
		_, err := m.DB.Exec(q)
		if err != nil {
			return ErrInsertDB
		}
	}
	return nil
}

func (m PgSQLRepository) AddContact(contact contact.Contact) error {
	panic("implement")
}

func (m PgSQLRepository) GetContact(id int) error {
	panic("implement")
}
