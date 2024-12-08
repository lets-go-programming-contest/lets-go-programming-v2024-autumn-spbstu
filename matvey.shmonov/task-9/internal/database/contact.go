package database

import (
	"errors"
	"regexp"
)

type Contact struct {
	ID    int
	Name  string
	Phone string
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
	re := regexp.MustCompile(`^\d{3}-\d{3}-\d{4}$`)
	if !re.MatchString(phone) {
		return errors.New("phone must be in the format 123-456-7890")
	}
	return nil
}
