package commonErrors

import "errors"

var (
	ErrContactNotExists     = errors.New("contact does not exist")
	ErrContactAlreadyExists = errors.New("contact already exists")
	ErrInternalError        = errors.New("internal error")
	ErrIncorrectNumber      = errors.New("incorrect number")
)
