package service

import (
	"context"
	"fmt"

	"github.com/EmptyInsid/task-9/internal/models"
)

type database interface {
	GetContacts(ctx context.Context) ([]models.Contact, error)
	GetContact(ctx context.Context, id int) (*models.Contact, error)
	CreateContact(ctx context.Context, newContact models.Contact) (int, error)
	UpdateContact(ctx context.Context, contact models.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

type DbService struct {
	db database
}

func NewDbService(db database) DbService {
	return DbService{
		db: db,
	}
}

func (s *DbService) GetContacts() ([]models.Contact, error) {
	contacts, err := s.db.GetContacts(context.Background())
	if err != nil {
		return nil, fmt.Errorf("service :: %w; db :: %w", errGetContacts, err)
	}
	return contacts, nil
}

func (s *DbService) GetContact(id int) (*models.Contact, error) {
	contact, err := s.db.GetContact(context.Background(), id)
	if err != nil {
		return nil, fmt.Errorf("service :: %w; db :: %w", errGetContact, err)
	}
	return contact, nil
}

func (s *DbService) CreateContact(contact models.Contact) (int, error) {
	id, err := s.db.CreateContact(context.Background(), contact)
	if err != nil {
		return 0, fmt.Errorf("service :: %w; db :: %w", errCreateContact, err)
	}
	return id, nil
}

func (s *DbService) UpdateContact(contact models.Contact) error {
	if err := s.db.UpdateContact(context.Background(), contact); err != nil {
		return fmt.Errorf("service :: %w; db :: %w", errUpdContact, err)
	}
	return nil
}

func (s *DbService) DeleteContact(id int) error {
	if err := s.db.DeleteContact(context.Background(), id); err != nil {
		return fmt.Errorf("service :: %w; db :: %w", errDelContact, err)
	}
	return nil
}
