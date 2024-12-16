package services

import (
	"github.com/nayzzerr/task-9/internal/models"
	"github.com/nayzzerr/task-9/internal/repository"
)

type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) GetAllContacts() ([]models.Contact, error) {
	return s.repo.GetAll()
}

func (s *ContactService) GetContactByID(id int) (*models.Contact, error) {
	return s.repo.GetByID(id)
}

func (s *ContactService) CreateContact(contact *models.Contact) error {
	return s.repo.Create(contact)
}

func (s *ContactService) UpdateContact(contact *models.Contact) error {
	return s.repo.Update(contact)
}

func (s *ContactService) DeleteContact(id int) error {
	return s.repo.Delete(id)
}
