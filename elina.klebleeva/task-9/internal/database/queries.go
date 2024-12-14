package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	"github.com/EmptyInsid/task-9/internal/models"
)

func (db *Database) GetContacts(ctx context.Context) ([]models.Contact, error) {
	query := `SELECT * FROM contacts`

	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer rows.Close()

	contacts := []models.Contact{}

	for rows.Next() {
		var contact models.Contact

		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (db *Database) GetContact(ctx context.Context, id int) (*models.Contact, error) {
	query := `SELECT * FROM contacts WHERE id = $1`

	var contact models.Contact
	if err := db.pool.QueryRow(ctx, query, id).Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &contact, nil
}

func (db *Database) CreateContact(ctx context.Context, newContact models.Contact) (int, error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return 0, fmt.Errorf("%w", err)
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO contacts(name, phone) VALUES($1, $2) RETURNING id`

	var id int
	if err := tx.QueryRow(ctx, query, newContact.Name, newContact.Phone).Scan(&id); err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return 0, fmt.Errorf("%w", err)
	}

	return id, nil
}

func (db *Database) UpdateContact(ctx context.Context, contact models.Contact) (*models.Contact, error) {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	defer tx.Rollback(ctx)

	if contact.Name == "" {
		query := `UPDATE contacts SET phone = $1 WHERE id = $2 RETURNING id, name, phone`
		err = tx.QueryRow(ctx, query, contact.Phone, contact.ID).Scan(&contact.ID, &contact.Name, &contact.Phone)
	} else if contact.Phone == "" {
		query := `UPDATE contacts SET name = $1 WHERE id = $2 RETURNING id, name, phone`
		err = tx.QueryRow(ctx, query, contact.Name, contact.ID).Scan(&contact.ID, &contact.Name, &contact.Phone)
	} else {
		query := `UPDATE contacts SET name = $1, phone = $2 WHERE id = $3 RETURNING id, name, phone`
		err = tx.QueryRow(ctx, query, contact.Name, contact.Phone, contact.ID).Scan(&contact.ID, &contact.Name, &contact.Phone)
	}

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &contact, nil
}

func (db *Database) DeleteContact(ctx context.Context, id int) error {
	tx, err := db.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM contacts WHERE id = $1`

	commandTag, err := tx.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
