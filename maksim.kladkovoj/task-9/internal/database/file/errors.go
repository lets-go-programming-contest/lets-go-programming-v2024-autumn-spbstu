package database

import "errors"

var (
	ErrDatabaseConnect  = errors.New("database connect failed")
	ErrDatabasePing     = errors.New("database ping failed")
	ErrDatabaseQuery    = errors.New("database query failed")
	ErrScanContact      = errors.New("scanning contacts failed")
	ErrContactNotExists = errors.New("contact not exists")
	ErrPhoneInvalid     = errors.New("phone number is invalid")
	ErrGetID            = errors.New("getting id failed")
)
