package handlers

import "errors"

var (
	ErrCreate      = errors.New("error creating a contact")
	ErrUpdate      = errors.New("error updating the contact")
	ErrDelete      = errors.New("error deleting a contact")
	ErrGet         = errors.New("the contact was not found")
	ErrDecode      = errors.New("incorrect data")
	ErrPhoneFormat = errors.New("incorrect phone format: the number must start with '+7' and contain 10 digits")
)
