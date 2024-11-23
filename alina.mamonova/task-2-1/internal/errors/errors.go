package errors

import "errors"

var (
	InvalidOperator = errors.New("Invalid operator.")
	InvalidInput    = errors.New("Invalid input.")
)
