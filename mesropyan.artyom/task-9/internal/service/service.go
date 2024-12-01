package service

import (
	"errors"
	"fmt"

	"github.com/artem6554/task-9/internal/contacts"
)

type Service struct {
}

func (s Service) Upload(name string, number string) error {
	err := contacts.Exists(name)
	if err != nil {
		err = errors.New("contact already exists")
		return fmt.Errorf("%w: %s", err, name)
	}
	err = contacts.CorrectNumber(number)
	if err != nil {
		return fmt.Errorf("%w: %s", err, number)
	}
	err = contacts.AddContact(name, number)
	if err != nil {
		return fmt.Errorf("%w: %s, %s", err, name, number)
	}

	return nil
}

func (s Service) Delete(name string) error {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.New("contact does not exist")
		return fmt.Errorf("%w: %s", err, name)
	}

	err = contacts.DeleteContact(name)
	if err != nil {
		return fmt.Errorf("%w: %s", err, name)
	}
	return nil

}

func (s Service) Get(name string) ([]byte, error) {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.New("contact does not exist")
		return nil, fmt.Errorf("%w: %s", err, name)
	}

	data, err := contacts.GetContact(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, name)
	}
	return data, nil
}

func (s Service) Update(name string, number string) error {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.New("contact does not exist")
		return fmt.Errorf("%w: %s", err, name)
	}

	err = contacts.CorrectNumber(number)
	if err != nil {
		return fmt.Errorf("%w: %s", err, number)
	}

	err = contacts.EditNumber(name, number)
	if err != nil {
		return fmt.Errorf("%w: %s, %s", err, name, number)
	}
	return nil
}
