package service

import (
	"errors"

	"github.com/artem6554/task-9/internal/contacts"
)

type Service struct {
}

func (s Service) Upload(name string, number string) error {
	err := contacts.Exists(name)
	if err != nil {
		return err // TODO: wrap error nicely
	}
	err = contacts.CorrectNumber(number)
	if err != nil {
		return err // TODO: wrap error nicely
	}
	err = contacts.AddContact(name, number)
	if err != nil {
		return err // TODO: wrap error nicely
	}

	return nil
}

func (s Service) Delete(name string) error {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.New("contact does not exist")
		return err // TODO: wrap error nicely
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
		err = errors.New("contact does not exist")
		return nil, err // TODO: wrap error nicely
	}

	data, err := contacts.GetContact(name)
	if err != nil {
		return nil, err // TODO: wrap error nicely
	}
	return data, nil
}

func (s Service) Update(name string, number string) error {
	err := contacts.Exists(name)
	if err == nil {
		err = errors.New("contact does not exist")
		return err // TODO: wrap error nicely
	}

	err = contacts.CorrectNumber(number)
	if err != nil {
		return err // TODO: wrap error nicely
	}

	err = contacts.EditNumber(name, number)
	if err != nil {
		return err // TODO: wrap error nicely
	}
	return nil
}
