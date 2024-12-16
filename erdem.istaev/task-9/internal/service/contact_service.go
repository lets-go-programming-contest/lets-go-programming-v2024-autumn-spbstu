package service

import (
	"erdem.istaev/task-9/internal/repository/database"
	"erdem.istaev/task-9/internal/structure"
)

type ContactService struct {
	repository database.Contacts
}

func NewContactService(repository database.Contacts) *ContactService {
	return &ContactService{repository: repository}
}

func (s *ContactService) CreateContact(contact structure.Contact) (int, error) {
	if err := contact.IsValidPhone(); err != nil {
		return contact.Id, err
	}

	return s.repository.CreateContact(contact)
}

func (s *ContactService) UpdateContact(id int, contact structure.Contact) error {
	if err := contact.IsValidPhone(); err != nil {
		return err
	}

	return s.repository.UpdateContact(id, contact)
}

func (s *ContactService) DeleteContact(id int) error {
	return s.repository.DeleteContact(id)
}

func (s *ContactService) GetContactById(id int) (structure.Contact, error) {
	return s.repository.GetContactById(id)
}

func (s *ContactService) GetAllContacts() ([]structure.Contact, error) {
	return s.repository.GetAllContacts()
}
