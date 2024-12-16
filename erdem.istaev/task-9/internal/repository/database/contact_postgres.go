package database

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"

	"erdem.istaev/task-9/internal/structure"
)

type ContactPostgres struct {
	db *sqlx.DB
}

func NewContactPostgres(db *sqlx.DB) *ContactPostgres {
	return &ContactPostgres{db: db}
}

func (r *ContactPostgres) CreateContact(contact structure.Contact) (int, error) {
	var id int
	createContactsQuery := "INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id"
	row := r.db.QueryRow(createContactsQuery, contact.Name, contact.Phone)
	if err := row.Scan(&id); err != nil {
		return id, fmt.Errorf("error create contact with id %d: %w", id, err)
	}

	return id, nil
}

func (r *ContactPostgres) GetAllContacts() ([]structure.Contact, error) {
	var contacts []structure.Contact

	query := "SELECT * FROM contacts"
	err := r.db.Select(&contacts, query)

	return contacts, err
}

func (r *ContactPostgres) GetContactById(id int) (structure.Contact, error) {
	var contact structure.Contact

	query := "SELECT * FROM contacts WHERE id = $1"
	if err := r.db.Get(&contact, query, id); err != nil {
		return contact, fmt.Errorf("contact with id %d not found: %w", id, err)
	}

	return contact, nil
}

func (r *ContactPostgres) DeleteContact(id int) error {
	_, err := r.GetContactById(id)
	if err != nil {
		return err
	}

	deleteQuery := "DELETE FROM contacts WHERE id = $1"
	_, err = r.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("delete contact error: %w", err)
	}

	return nil
}

func (r *ContactPostgres) UpdateContact(id int, contact structure.Contact) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	_, err := r.GetContactById(id)
	if err != nil {
		return err
	}

	if contact.Name != "" {
		setValues = append(setValues, fmt.Sprintf("name = $%d", argId))
		args = append(args, contact.Name)
		argId++
	}

	if contact.Phone != "" {
		setValues = append(setValues, fmt.Sprintf("phone = $%d", argId))
		args = append(args, contact.Phone)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf("UPDATE contacts SET %s WHERE id = $%d ", setQuery, argId)
	args = append(args, id)

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("update contact error: %w", err)
	}

	return nil
}
