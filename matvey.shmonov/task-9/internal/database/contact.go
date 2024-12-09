package database

import (
	"errors"
	"regexp"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (c *Contact) Validate() error {
	if err := validateName(c.Name); err != nil {
		return err
	}
	if err := validatePhone(c.Phone); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(name) {
		return errors.New("name can only contain letters and spaces")
	}
	return nil
}

func validatePhone(phone string) error {
	re := regexp.MustCompile(`^\d{11}$`)
	if !re.MatchString(phone) {
		return errors.New("phone must be in the format 01234567890")
	}
	return nil
}
