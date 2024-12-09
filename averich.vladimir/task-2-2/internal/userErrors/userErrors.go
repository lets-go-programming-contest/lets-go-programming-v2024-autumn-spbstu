package userErrors

import "errors"

var (
	ErrorInvalidInput = errors.New("invalid input")
	ErrorOverflow     = errors.New("overflow")
	ErrorInvalidValue = errors.New("invalid value")
)
