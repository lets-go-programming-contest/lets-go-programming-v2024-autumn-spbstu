package io

import "fmt"

type ParseIntError struct {
	Actual string
}

func (err ParseIntError) Error() string {
	return fmt.Sprintf("ожидалось целое число. Получено: %v", err.Actual)
}
