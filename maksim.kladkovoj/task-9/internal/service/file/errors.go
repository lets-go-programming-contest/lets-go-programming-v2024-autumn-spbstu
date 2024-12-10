package service

import "errors"

var (
	ErrPost          = errors.New("create contact failed")
	ErrNotExists     = errors.New("contact not exists")
	ErrAlreadyExists = errors.New("contact already exists")
	ErrPut           = errors.New("update contact failed")
	ErrDelete        = errors.New("delete contact failed")
	ErrGet           = errors.New("getting contaccts failed")
)
