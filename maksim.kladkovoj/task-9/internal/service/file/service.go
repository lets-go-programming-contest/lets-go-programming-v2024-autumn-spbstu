package service

import (
	"fmt"

	database "github.com/Mmmakskl/task-9/internal/database/file"
)

type contactReadDB interface {
	Exists(string) error
	Get(string) ([]database.Contact, error)
}

type contactWriteDB interface {
	Post(database.Contact) error
	Put(database.Contact) error
	Delete(int) error
}

type service struct {
	readDB  contactReadDB
	writeDB contactWriteDB
}

func NewService(read contactReadDB, write contactWriteDB) service {
	return service{
		readDB:  read,
		writeDB: write,
	}
}

func (s *service) Post(name string, phone string) error {
	contact := database.Contact{
		Name:  name,
		Phone: phone,
	}

	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM contacts WHERE name = '%s' AND phone = '%s')", name, phone)
	if existsErr := s.readDB.Exists(query); existsErr == nil {
		return fmt.Errorf("%w: %w", ErrAlreadyExists, existsErr)
	}

	postErr := s.writeDB.Post(contact)
	if postErr != nil {
		return fmt.Errorf("%w: %w", ErrPost, postErr)
	}

	return nil
}

func (s *service) Put(id int, name string, phone string) error {
	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM contacts WHERE id = %d)", id)
	if existsErr := s.readDB.Exists(query); existsErr != nil {
		return fmt.Errorf("%w: %w: id:%d", ErrNotExists, existsErr, id)
	}

	contact := database.Contact{
		ID:    id,
		Name:  name,
		Phone: phone,
	}

	putErr := s.writeDB.Put(contact)
	if putErr != nil {
		return fmt.Errorf("%w: %w", ErrPut, putErr)
	}

	return nil
}

func (s *service) Delete(id int) error {
	query := fmt.Sprintf("SELECT EXISTS(SELECT * FROM contacts WHERE id = %d)", id)
	if existsErr := s.readDB.Exists(query); existsErr != nil {
		return fmt.Errorf("%w: %w: id:%d", ErrNotExists, existsErr, id)
	}

	deleteErr := s.writeDB.Delete(id)
	if deleteErr != nil {
		return fmt.Errorf("%w: %w", ErrDelete, deleteErr)
	}

	return nil
}

func (s *service) GetAll() ([]database.Contact, error) {
	if existsErr := s.readDB.Exists("SELECT EXISTS(SELECT id, name, phone FROM contacts)"); existsErr != nil {
		return nil, fmt.Errorf("%w: %w", ErrNotExists, existsErr)
	}

	contacts, getErr := s.readDB.Get("SELECT id, name, phone FROM contacts;")
	if getErr != nil {
		return nil, fmt.Errorf("%w: %w", ErrGet, getErr)
	}

	return contacts, nil
}

func (s *service) GetID(id int) ([]database.Contact, error) {
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM contacts WHERE id = %d)", id)
	if existsErr := s.readDB.Exists(query); existsErr != nil {
		return nil, fmt.Errorf("%w: %w", ErrNotExists, existsErr)
	}

	queryGet := fmt.Sprintf("SELECT id, name, phone FROM contacts WHERE id = %d", id)

	contacts, getErr := s.readDB.Get(queryGet)
	if getErr != nil {
		return nil, fmt.Errorf("%w: %w", ErrGet, getErr)
	}

	return contacts, nil
}
