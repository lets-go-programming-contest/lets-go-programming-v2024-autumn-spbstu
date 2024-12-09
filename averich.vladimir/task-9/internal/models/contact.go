package models

import (
	"errors"
	"regexp"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ValidateContact(contact *Contact) error {

	if contact.Name == "" {
		return errors.New("Name is required")
	}

	if contact.Phone == "" {
		return errors.New("Phone number is required")
	}

	re := regexp.MustCompile(`^\+?[0-9]{10,14}$`)
	if !re.MatchString(contact.Phone) {
		return errors.New("Invalid phone number format")
	}
	return nil
}
