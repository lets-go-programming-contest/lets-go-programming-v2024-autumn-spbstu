package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hahapathetic/task-9/internal/commonErrors"
	"github.com/hahapathetic/task-9/internal/contacts"
)

type Service struct {
}

func (s Service) Upload(name string, number string) error {
	err := contacts.ExistsByNumber(number)
	if err != nil {
		if errors.Is(err, commonErrors.ErrContactAlreadyExists) {
			return fmt.Errorf("%w: %q", err, number)
		}
		return err
	}

	err = contacts.IsCorrectNumber(number)
	if err != nil {
		return fmt.Errorf("%w: %q", err, number)
	}

	err = contacts.AddContact(name, number)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Delete(id string) error {

	err := contacts.ExistsByID(id)
	if err != nil {
		if errors.Is(err, commonErrors.ErrContactNotExists) {
			return fmt.Errorf("%w: %q", err, id)
		}
		return err
	}

	err = contacts.DeleteContactByID(id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Get(id string) ([]byte, error) {

	contact, err := contacts.GetContactByID(id)
	if err != nil {
		if errors.Is(err, commonErrors.ErrContactNotExists) {
			return nil, fmt.Errorf("%w: %q", err, id)
		}
		return nil, err
	}

	data, err := json.Marshal(contact)
	if err != nil {
		return nil, fmt.Errorf("marshal contact: %w", err)
	}

	return data, nil
}

func (s Service) Update(id string, name string, number string) error {

	err := contacts.ExistsByID(id)
	if err != nil {
		if errors.Is(err, commonErrors.ErrContactNotExists) {
			return fmt.Errorf("%w: %q", err, id)
		}
		return err
	}

	currentContact, err := contacts.GetContactByID(id)
	if err != nil {
		return err
	}

	if number != currentContact.Number {
		err = contacts.ExistsByNumber(number)
		if err != nil {
			if errors.Is(err, commonErrors.ErrContactAlreadyExists) {
				return fmt.Errorf("%w: %q", err, number)
			}
			return err
		}
	}

	err = contacts.IsCorrectNumber(number)
	if err != nil {
		return fmt.Errorf("%w: %q", err, number)
	}

	err = contacts.EditNumberByID(id, name, number)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) GetAll() ([]byte, error) {
	data, err := contacts.GetAllContacts()
	if err != nil {
		return nil, err
	}
	return data, nil
}
