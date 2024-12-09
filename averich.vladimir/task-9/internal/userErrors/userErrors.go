package userErrors

import "errors"

var (
	ErrInitDB = errors.New("error initialization of DB")
)
