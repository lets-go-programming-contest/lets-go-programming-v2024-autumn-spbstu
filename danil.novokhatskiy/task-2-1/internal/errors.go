package internal

import "fmt"

type ErrorTemp struct{}

func (errorTemp ErrorTemp) Error() string {
	return fmt.Sprintf("-1")
}
