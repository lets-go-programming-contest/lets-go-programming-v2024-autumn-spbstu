package errorUtils

import (
	"fmt"
	"runtime"
)

func ErrorWithLocation(err error) error {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		return fmt.Errorf("%s:%d:\n%w", file, line, err)
	}
	return err
}
