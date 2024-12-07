package service

import (
	"fmt"

	errs "errors"

	"github.com/artem6554/task-9/internal/errors"

	"github.com/artem6554/task-9/internal/contacts"
)

type Service struct {
}

func (s Service) Upload(name string, number string) error {
	err := contacts.Exists(name)
	if err != nil {
		if errs.Is(err, errors.ErrContactAlreadyExists) {
			return fmt.Errorf("%w: %q", err, name)
		}
		return err
	}
	err = contacts.CorrectNumber(number)
	if err != nil {
		return fmt.Errorf("%w: %q", err, number)
	}
	err = contacts.AddContact(name, number)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Delete(name string) error {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.ErrContacteNotExists
		return fmt.Errorf("%w: %q", err, name)
	}

	err = contacts.DeleteContact(name)
	if err != nil {
		return err
	}
	return nil

}

func (s Service) Get(name string) ([]byte, error) {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.ErrContacteNotExists
		return nil, fmt.Errorf("%w: %q", err, name)
	}

	data, err := contacts.GetContact(name)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s Service) Update(name string, number string) error {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.ErrContacteNotExists
		return fmt.Errorf("%w: %q", err, name)
	}

	err = contacts.CorrectNumber(number)
	if err != nil {
		return fmt.Errorf("%w: %q", err, number)
	}

	err = contacts.EditNumber(name, number)
	if err != nil {
		return err
	}
	return nil
}
