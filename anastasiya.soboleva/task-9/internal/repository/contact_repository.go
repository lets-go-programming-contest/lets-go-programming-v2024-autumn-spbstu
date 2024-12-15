package repository

import (
	"database/sql"
	"errors"

	"github.com/nayzzerr/task-9/internal/models"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) GetAll() ([]models.Contact, error) {
	rows, err := r.db.Query("SELECT id, name, phone FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func (r *ContactRepository) GetByID(id int) (*models.Contact, error) {
	var contact models.Contact
	err := r.db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id).
		Scan(&contact.ID, &contact.Name, &contact.Phone)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return &contact, err
}

func (r *ContactRepository) Create(contact *models.Contact) error {
	return r.db.QueryRow(
		"INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id",
		contact.Name, contact.Phone).Scan(&contact.ID)
}

func (r *ContactRepository) Update(contact *models.Contact) error {
	_, err := r.db.Exec("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3",
		contact.Name, contact.Phone, contact.ID)
	return err
}

func (r *ContactRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM contacts WHERE id = $1", id)
	return err
}
