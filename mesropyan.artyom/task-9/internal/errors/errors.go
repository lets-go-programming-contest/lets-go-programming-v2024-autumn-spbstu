package errors

import "errors"

var (
	// ErrFileNotExists using when contact does not exist in db
	ErrContacteNotExists = errors.New("contact does not exist")
	// ErrFileAlreadyExists using when contact already exists in db
	ErrContactAlreadyExists = errors.New("contact alredy exists")
	// ErrInternalError using if something went wrong
	ErrInternalError = errors.New("internal error")
	// ErrIncorrectNumber using when number from user is wrong
	ErrIncorrectNumber = errors.New("incorrect number")
)
