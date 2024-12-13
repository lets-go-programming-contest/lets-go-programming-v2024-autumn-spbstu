package errors

import (
	"errors"
)

var (
	ErrNoContact        = errors.New("no such contact in list")
	ErrEmptyData        = errors.New("name or phome shouldn't be empty")
	ErrExistContact     = errors.New("contact with such phone number already exists")
	ErrWrongPhoneFormat = errors.New("incorrect phone format, use like: \n\t+79261234567\n\t8(926)123-45-67\n\t123-45-67\n\t9261234567")
	ErrEncodeJson       = errors.New("error while encoding contact")
	ErrDecodeJson       = errors.New("error while decode contact")
	ErrInternal         = errors.New("some internal err")
)
