package errors

import (
	"errors"
)

var (
	ErrNoContact        = errors.New("no such contact in list")
	ErrEmptyData        = errors.New("name or phome shouldn't be empty")
	ErrExistContact     = errors.New("contact with such phone number already exists")
	ErrWrongPhoneFormat = errors.New("incorrect phone format, use like: 8 (800) 555-35-55")
	ErrEncodeJSON       = errors.New("error while encoding contact")
	ErrDecodeJSON       = errors.New("error while decode contact")
	ErrInternal         = errors.New("some internal err")
)
