package database

import (
	"github.com/jmoiron/sqlx"

	"erdem.istaev/task-9/internal/structure"
)

type Contacts interface {
	CreateContact(contact structure.Contact) (int, error)
	GetAllContacts() ([]structure.Contact, error)
	GetContactById(id int) (structure.Contact, error)
	DeleteContact(id int) error
	UpdateContact(id int, contact structure.Contact) error
}

type Repository struct {
	Contacts
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Contacts: NewContactPostgres(db),
	}
}
