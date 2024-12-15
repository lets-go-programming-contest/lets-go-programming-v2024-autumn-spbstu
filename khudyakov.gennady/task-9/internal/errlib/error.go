package errlib

import (
	"errors"
	"net/http"
)

var ErrInternal = errors.New("some internal error happened")

var ErrBadRequest = errors.New("bad request")

var ErrResourceAlreadyExists = errors.New("resource already exists")

var ErrNotFound = errors.New("resource not found")

type JSONError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"msg"`
	} `json:"error"`
}

func GetJSONError(err error) *JSONError {
	var jsonErr = JSONError{}
	jsonErr.Error.Message = err.Error()
	switch {
	case errors.Is(err, ErrInternal):
		jsonErr.Error.Code = http.StatusInternalServerError
	case errors.Is(err, ErrResourceAlreadyExists):
		jsonErr.Error.Code = http.StatusConflict
	case errors.Is(err, ErrNotFound):
		jsonErr.Error.Code = http.StatusNotFound
	case errors.Is(err, ErrBadRequest):
		jsonErr.Error.Code = http.StatusBadRequest
	default:
		jsonErr.Error.Message = "some internal error happened"
		jsonErr.Error.Code = http.StatusInternalServerError
	}

	return &jsonErr
}
