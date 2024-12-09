package userErrors

import "errors"

var (
	ErrOverflow       = errors.New("overflow")
	ErrIncorrectInput = errors.New("incorrect input")
)
