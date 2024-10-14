package io

import "fmt"

func ParseIntError(actual string) error {
	return fmt.Errorf("ожидалось целое число. Получено: %v", actual)
}
