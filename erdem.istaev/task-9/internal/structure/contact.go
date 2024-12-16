package structure

import (
	"errors"
	"regexp"
	"strings"
)

type Contact struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var (
	ErrValidPhone = errors.New("invalid phone number")
)

func (c *Contact) IsValidPhone() error {
	phone := c.Phone
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")

	re := regexp.MustCompile(`^\+?[7-8]?\d{10}$`)
	if !re.MatchString(phone) {
		return ErrValidPhone
	}

	return nil
}
