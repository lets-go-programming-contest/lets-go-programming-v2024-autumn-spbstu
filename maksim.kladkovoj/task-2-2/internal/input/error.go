package input

import "errors"

var (
	ErrInput error = errors.New("Input error")
	ErrK_th  error = errors.New("The number of dishes is less than the k-th number")
)
