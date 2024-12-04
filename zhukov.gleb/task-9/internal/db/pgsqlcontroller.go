package db

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"task-9/internal/contact"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var (
	ErrOpenDB    = errors.New("error open")
	ErrPingDB    = errors.New("error ping")
	ErrInsertDB  = errors.New("error insert")
	ErrNoContact = errors.New("contact not found")
)

type PgSQLRepository struct {
	DB *sql.DB
}

// TODO структуру передавать вместо полей
func NewPgSQLController(config Cfg) (PgSQLRepository, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.UDBName, config.UDBPass, config.PgSQLHost, config.PortPgSQL, config.DBPgSQLName)

	repo := &PgSQLRepository{}
	var err error

	repo.DB, err = sql.Open("postgres", dsn)
	if err != nil {
		return *repo, fmt.Errorf("%w + pgsql; %w", ErrOpenDB, err)
	}

	err = repo.DB.Ping()
	if err != nil {
		err_ := repo.DB.Close()
		if err_ != nil {
			return PgSQLRepository{}, errors.Join(err, err_)
		}
		return *repo, fmt.Errorf("%w + pgsql; %w", ErrPingDB, err)
	}

	err = repo.newPostgresStorage()
	if err != nil {
		err_ := repo.DB.Close()
		if err_ != nil {
			return PgSQLRepository{}, errors.Join(err, err_)
		}
		return *repo, err
	}

	return *repo, nil
}

func (m *PgSQLRepository) newPostgresStorage() error {
	driver, err := pgx.WithInstance(m.DB, &pgx.Config{})
	if err != nil {
		return err
	}

	migr, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		return err
	}

	if err = migr.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func (m *PgSQLRepository) AddContact(name, phone string, time time.Time) (contact.Contact, error) {
	q := "INSERT INTO contacts (Name, Phone, CreatedAt, UpdatedAt) VALUES($1, $2, $3, $4) RETURNING ID, Name, Phone, CreatedAt, UpdatedAt"

	var newContact contact.Contact

	err := m.DB.QueryRow(q, name, phone, time, time).Scan(&newContact.ID, &newContact.Name, &newContact.Phone, &newContact.CreatedAt, &newContact.UpdatedAt)
	if err != nil {
		return contact.Contact{}, fmt.Errorf("%w add user: %w", ErrInsertDB, err)
	}

	return newContact, nil
}

func (m *PgSQLRepository) GetContact(id int) (contact.Contact, error) {
	q := "SELECT ID, Name, Phone, CreatedAt, UpdatedAt FROM contacts WHERE ID = $1"

	var newContact contact.Contact

	err := m.DB.QueryRow(q, id).Scan(
		&newContact.ID,
		&newContact.Name,
		&newContact.Phone,
		&newContact.CreatedAt,
		&newContact.UpdatedAt,
	)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return contact.Contact{}, fmt.Errorf("%w getContact DB: %d", ErrNoContact, id)
	default:
		return contact.Contact{}, fmt.Errorf("getContact DB err: %w", err)
	}

	return newContact, nil
}

func (m *PgSQLRepository) GetAllContacts() ([]contact.Contact, error) {
	q := "SELECT ID, Name, Phone, CreatedAt, UpdatedAt FROM contacts"

	rows, err := m.DB.Query(q)
	if err != nil {
		return nil, fmt.Errorf("GetAllContacts DB: %w", err)
	}
	defer rows.Close()

	contacts := make([]contact.Contact, 0)

	for rows.Next() {
		var newContact contact.Contact
		err = rows.Scan(
			&newContact.ID,
			&newContact.Name,
			&newContact.Phone,
			&newContact.CreatedAt,
			&newContact.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("GetAllContacts DB: %w", err)
		}
		contacts = append(contacts, newContact)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllContacts DB: %w", err)
	}

	return contacts, nil
}

func (m *PgSQLRepository) UpdateContact(id int, name, phone string, time time.Time) (contact.Contact, error) {
	q := "UPDATE contacts SET Name = $1, Phone = $2, UpdatedAt = $3 WHERE ID = $4 RETURNING ID, Name, Phone, CreatedAt, UpdatedAt"

	var updatedContact contact.Contact

	err := m.DB.QueryRow(q, name, phone, time, id).Scan(
		&updatedContact.ID,
		&updatedContact.Name,
		&updatedContact.Phone,
		&updatedContact.CreatedAt,
		&updatedContact.UpdatedAt,
	)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return contact.Contact{}, fmt.Errorf("%w UpdateContact DB: %d", ErrNoContact, id)
	default:
		return contact.Contact{}, fmt.Errorf("UpdateContact DB err: %w", err)
	}

	return updatedContact, nil
}

func (m *PgSQLRepository) DeleteContact(id int) error {
	q := "DELETE FROM contacts WHERE ID = $1"

	result, err := m.DB.Exec(q, id)
	if err != nil {
		return fmt.Errorf("RowsAffected DB failed to execute query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected DB err: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%w getContact: %d", ErrNoContact, id)
	}

	return nil
}
