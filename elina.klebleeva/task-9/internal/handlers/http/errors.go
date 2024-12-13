package handlers

import (
	"errors"
	"net/http"

	myErr "github.com/EmptyInsid/task-9/internal/errors"
)

// errorMap is map with errors for http handle func.
type errorMap map[error]int

var (
	errorsGetMap = errorMap{
		myErr.ErrNoContact:  http.StatusNotFound,
		myErr.ErrEncodeJson: http.StatusInternalServerError,
		myErr.ErrInternal:   http.StatusInternalServerError,
	}

	errorsCreateMap = errorMap{
		myErr.ErrExistContact:     http.StatusConflict,
		myErr.ErrEmptyData:        http.StatusBadRequest,
		myErr.ErrWrongPhoneFormat: http.StatusBadRequest,
		myErr.ErrDecodeJson:       http.StatusBadRequest,
		myErr.ErrEncodeJson:       http.StatusInternalServerError,
		myErr.ErrInternal:         http.StatusInternalServerError,
	}

	errorsUpdMap = errorMap{
		myErr.ErrNoContact:        http.StatusNotFound,
		myErr.ErrEmptyData:        http.StatusBadRequest,
		myErr.ErrWrongPhoneFormat: http.StatusBadRequest,
		myErr.ErrDecodeJson:       http.StatusBadRequest,
		myErr.ErrEncodeJson:       http.StatusInternalServerError,
		myErr.ErrInternal:         http.StatusInternalServerError,
	}

	errorsDeleteMap = errorMap{
		myErr.ErrNoContact:  http.StatusNotFound,
		myErr.ErrEncodeJson: http.StatusInternalServerError,
		myErr.ErrInternal:   http.StatusInternalServerError,
	}
)

// getMappedStatusCode return code with errorMap and err.
// If err is nil returns http.StatusOK.
// If err can't be find in errorMap returns http.StatusInternalServerError.
func getStatusCode(m errorMap, err error) int {
	if err == nil {
		return http.StatusOK
	}
	for e, c := range m {
		if errors.Is(err, e) {
			return c
		}
	}
	return http.StatusInternalServerError
}
