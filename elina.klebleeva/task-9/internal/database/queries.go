package database

import (
	"context"
	"fmt"

	"github.com/EmptyInsid/task-9/internal/models"
)

func (db *Database) GetContacts(ctx context.Context) ([]models.Contact, error) {
	query := `SELECT * FROM contacts`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	contacts := []models.Contact{}
	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.Id, &contact.Name, &contact.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (db *Database) GetContact(ctx context.Context, name string) (*models.Contact, error) {
	query := `SELECT * FROM contacts WHERE name = $1`

	var contact models.Contact
	if err := db.pool.QueryRow(ctx, query, name).Scan(&contact.Id, &contact.Name, &contact.Phone); err != nil {
		return nil, err
	}

	return &contact, nil
}

func (db *Database) CreateContact(ctx context.Context, newContact models.Contact) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO contacts(name, phone) VALUES($1, $2)`

	commandTag, err := tx.Exec(ctx, query, newContact.Name, newContact.Phone)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("empty row after insert contact")
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}
	return nil

}

func (db *Database) UpdateContact(ctx context.Context, newName, oldName string) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return nil
	}
	defer tx.Rollback(ctx)

	query := `UPDATE contacts SET name = $1 WHERE name = $2`

	commandTag, err := tx.Exec(ctx, query, newName, oldName)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("empty row after update contact")
	}

	if err := tx.Commit(ctx); err != nil {
		return nil
	}

	return nil
}

func (db *Database) DeleteContact(ctx context.Context, contact models.Contact) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM contacts WHERE name = $1 AND phone = $2`

	commandTag, err := tx.Exec(ctx, query, contact.Name, contact.Phone)
	if err != nil {
		return nil
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("empty row after delete contact")
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil

}
