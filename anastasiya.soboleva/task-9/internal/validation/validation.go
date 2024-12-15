package validation

import (
	"errors"
	"regexp"
	"strings"
)

func ValidatePhoneNumber(phone string) error {
	phoneRegex := regexp.MustCompile(`^(\+7|7|8)?\d{10}$`)
	if !phoneRegex.MatchString(phone) {
		return errors.New("invalid phone number format")
	}
	return nil
}

func ValidateName(name string) error {
	name = strings.TrimSpace(name)

	if len(name) < 2 || len(name) > 100 {
		return errors.New("name must be between 2 and 100 characters")
	}

	nameRegex := regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ]+([-\s][a-zA-Zа-яА-ЯёЁ]+)*$`)
	if !nameRegex.MatchString(name) {
		return errors.New("name contains invalid characters")
	}

	return nil
}

func ValidateContactData(name, phone string) error {
	if err := ValidateName(name); err != nil {
		return err
	}
	if err := ValidatePhoneNumber(phone); err != nil {
		return err
	}
	return nil
}
