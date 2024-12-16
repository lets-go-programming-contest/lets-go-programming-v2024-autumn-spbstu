package userErrors

import "errors"

var (
	ErrorInvalidInput = errors.New("invalid input")
	ErrorOverflow     = errors.New("overflow")
	ErrorInvalidValue = errors.New("invalid value")
	ErrorOutOfRange   = errors.New("the index is out of range")
)
