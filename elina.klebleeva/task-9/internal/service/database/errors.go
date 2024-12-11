package service

import "errors"

var (
	errGetContacts   = errors.New("error while get contacts")
	errGetContact    = errors.New("error while get contact")
	errCreateContact = errors.New("error while create contact")
	errUpdContact    = errors.New("error while update contact")
	errDelContact    = errors.New("error while delete contact")
)
