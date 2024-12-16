package handler

import (
	"errors"
	"net/http"

	MyError "github.com/hahapathetic/task-9/internal/commonErrors"
)

type errorMap map[error]int

var (
	errorsGetMap = errorMap{
		MyError.ErrContactNotExists: http.StatusNotFound,
		MyError.ErrInternalError:    http.StatusInternalServerError,
	}

	errorsUpdateMap = errorMap{
		MyError.ErrContactNotExists: http.StatusNotFound,
		MyError.ErrIncorrectNumber:  http.StatusNotFound,
		MyError.ErrInternalError:    http.StatusInternalServerError,
	}

	errorsUploadMap = errorMap{
		MyError.ErrIncorrectNumber:      http.StatusNotFound,
		MyError.ErrContactAlreadyExists: http.StatusConflict,
		MyError.ErrInternalError:        http.StatusInternalServerError,
	}

	errorsDeleteMap = errorMap{
		MyError.ErrContactNotExists: http.StatusNotFound,
		MyError.ErrInternalError:    http.StatusInternalServerError,
	}
)

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
