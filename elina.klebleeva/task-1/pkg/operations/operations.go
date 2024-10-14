package operations

import (
	"errors"
)

func Add(a float64, b float64) float64 {
	return a + b
}

func Multiply(a float64, b float64) float64 {
	return a * b
}

func Subtract(a float64, b float64) float64 {
	return a - b
}

func Divide(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("делить на ноль нельзя")
	}
	return a / b, nil
}
