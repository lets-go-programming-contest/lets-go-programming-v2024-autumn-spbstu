package input

import "errors"

var (
	ErrInput  error = errors.New("Input error")
	ErrTemp   error = errors.New("Error converting temperature")
	ErrRegexp error = errors.New("Error converting regexp string")
)
