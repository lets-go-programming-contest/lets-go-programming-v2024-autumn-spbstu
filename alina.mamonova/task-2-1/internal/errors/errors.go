package errors

import "errors"

var (
	InvalidOperator = errors.New("Invalid operator. Please enter one of the following: <=, >=.")
	InvalidInput    = errors.New("Invalid input. Please enter a valid number.")
)
