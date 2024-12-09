package http

import (
	"errors"
	"net/http"

	MyError "github.com/artem6554/task-9/internal/errors"
)

// errorMap is map with errors for http handle func.
type errorMap map[error]int

var (
	errorsGetMap = errorMap{
		MyError.ErrContacteNotExists: http.StatusNotFound,
		MyError.ErrInternalError:     http.StatusInternalServerError,
	}

	errorsUpdateMap = errorMap{
		MyError.ErrContacteNotExists: http.StatusNotFound,
		MyError.ErrIncorrectNumber:   http.StatusNotFound,
		MyError.ErrInternalError:     http.StatusInternalServerError,
	}

	errorsUploadMap = errorMap{
		MyError.ErrIncorrectNumber:      http.StatusNotFound,
		MyError.ErrContactAlreadyExists: http.StatusConflict,
		MyError.ErrInternalError:        http.StatusInternalServerError,
	}

	errorsDeleteMap = errorMap{
		MyError.ErrContacteNotExists: http.StatusNotFound,
		MyError.ErrInternalError:     http.StatusInternalServerError,
	}
)

// getMappedStatusCode return code with errorMap and err.
// If err is nil returns http.StatusOK.
// If err can't be find in errorMap returns http.StatusInternalServerError.
func getMappedStatusCode(m errorMap, err error) int {
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
