package calculation

import (
	"errors"
	"task-1/pkg/operations"
)

func Calculation(a float64, b float64, operation string) (float64, error) {
	var result float64
	var err error

	switch operation {
	case "*":
		result = operations.Multiply(a, b)
	case "+":
		result = operations.Add(a, b)
	case "-":
		result = operations.Subtract(a, b)
	case "/":
		result, err = operations.Divide(a, b)
		if err != nil {
			return result, err
		}
	default:
		return 0, errors.New("некорректная операция, пожалуйста, используйте символы +, -, * или /")
	}

	return result, nil
}
