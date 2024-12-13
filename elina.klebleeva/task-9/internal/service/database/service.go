package service

import (
	"context"
	"errors"
	"fmt"

	myErr "github.com/EmptyInsid/task-9/internal/errors"
	"github.com/EmptyInsid/task-9/internal/models"
	"github.com/jackc/pgx/v5"
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("service :: %w :: db :: %w", myErr.ErrInternal, err)
	}
	return contacts, nil
}

func (s *DbService) GetContact(id int) (*models.Contact, error) {
	contact, err := s.db.GetContact(context.Background(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrNoContact, id, err)
		}
		return nil, fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrInternal, id, err)
	}
	return contact, nil
}

func (s *DbService) CreateContact(contact models.Contact) (int, error) {
	id, err := s.db.CreateContact(context.Background(), contact)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrNoContact, id, err)
		}
		return 0, fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrInternal, id, err)
	}
	return id, nil
}

func (s *DbService) UpdateContact(contact models.Contact) error {
	if err := s.db.UpdateContact(context.Background(), contact); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrNoContact, contact.Id, err)
		}
		return fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrInternal, contact.Id, err)
	}
	return nil
}

func (s *DbService) DeleteContact(id int) error {
	if err := s.db.DeleteContact(context.Background(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrNoContact, id, err)
		}
		return fmt.Errorf("[service] %w :: [id] %d :: [db] %w", myErr.ErrInternal, id, err)
	}
	return nil
}
