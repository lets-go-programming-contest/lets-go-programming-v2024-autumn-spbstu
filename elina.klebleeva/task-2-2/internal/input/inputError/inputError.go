package inputError

import "fmt"

type IncorrectNumber struct {
	ErrorPlace string
}

func (err *IncorrectNumber) Error() string {
	return fmt.Sprintf("incorrect numder input in: %s", err.ErrorPlace)
}
