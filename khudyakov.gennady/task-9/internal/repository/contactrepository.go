package repository

import (
	"context"
	"fmt"

	"github.com/KRYST4L614/task-9/internal/domain"
	"github.com/KRYST4L614/task-9/internal/errlib"
	"github.com/jackc/pgx/v4/pgxpool"
)

type contactRepository struct {
	dbPool *pgxpool.Pool
}

func NewRepository(dbPool *pgxpool.Pool) ContactRepositoryIface {
	return &contactRepository{
		dbPool: dbPool,
	}
}

func (cr *contactRepository) Get(ctx context.Context, id int) (*domain.Contact, error) {
	tx, err := cr.dbPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	var contact domain.Contact
	err = tx.QueryRow(ctx, `SELECT * FROM contacts WHERE ID = $1`, id).Scan(&contact.Id, &contact.Name, &contact.Phone)

	if contact.Id == 0 {
		return nil, fmt.Errorf("%w: contact with id = %v not found", errlib.ErrNotFound, id)
	}

	if err != nil {
		return nil, err
	}
	tx.Commit(ctx)
	return &contact, nil
}

func (cr *contactRepository) GetAll(ctx context.Context) ([]*domain.Contact, error) {
	tx, err := cr.dbPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	contacts := []*domain.Contact{}
	rows, err := tx.Query(ctx, `SELECT * FROM contacts`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var contact domain.Contact
		err := rows.Scan(&contact.Id, &contact.Name, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}

	tx.Commit(ctx)

	return contacts, nil
}

func (cr *contactRepository) Create(ctx context.Context, contact domain.Contact) (*domain.Contact, error) {
	tx, err := cr.dbPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	var createdContact domain.Contact
	err = tx.QueryRow(ctx, `INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING *`,
		contact.Name, contact.Phone).Scan(&createdContact.Id, &createdContact.Name, &createdContact.Phone)

	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	return &createdContact, nil
}

func (cr *contactRepository) Update(ctx context.Context, contact domain.Contact) (*domain.Contact, error) {
	tx, err := cr.dbPool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	var createdContact domain.Contact
	err = tx.QueryRow(ctx, `UPDATE contacts SET name = $1, phone = $2 WHERE id = $3 RETURNING *`,
		contact.Name, contact.Phone, contact.Id).Scan(&createdContact.Id, &createdContact.Name, &createdContact.Phone)

	if createdContact.Id == 0 {
		return nil, fmt.Errorf("%w: contact with id = %v not found", errlib.ErrNotFound, contact.Id)
	}

	if err != nil {
		return nil, err
	}

	tx.Commit(ctx)

	return &createdContact, nil
}

func (cr *contactRepository) DeleteById(ctx context.Context, id int) error {
	tx, err := cr.dbPool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `DELETE FROM contacts WHERE ID = $1`, id)
	if err != nil {
		return err
	}
	tx.Commit(ctx)
	return nil
}
