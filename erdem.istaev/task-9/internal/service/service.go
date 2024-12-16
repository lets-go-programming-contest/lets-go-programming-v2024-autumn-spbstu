package service

import (
	"erdem.istaev/task-9/internal/repository/database"
	"erdem.istaev/task-9/internal/structure"
)

type Contacts interface {
	CreateContact(contact structure.Contact) (int, error)
	GetAllContacts() ([]structure.Contact, error)
	GetContactById(id int) (structure.Contact, error)
	DeleteContact(id int) error
	UpdateContact(id int, contact structure.Contact) error
}

type Service struct {
	Contacts
}

func NewService(repository *database.Repository) *Service {
	return &Service{
		Contacts: NewContactService(repository.Contacts),
	}
}
