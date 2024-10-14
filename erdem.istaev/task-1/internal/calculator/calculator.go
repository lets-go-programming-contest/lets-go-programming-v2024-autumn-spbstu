package calculator

import (
	"errors"
)

var (
	ErrZeroDivision       = errors.New("Ошибка: деление на ноль невозможно.")
	ErrIncorrectOperation = errors.New("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
)

func Calculate(op1, op2 float64, operation string) (float64, error) {
	switch operation {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return 0, ErrZeroDivision
		} else {
			return op1 / op2, nil
		}
	default:
		return 0, ErrIncorrectOperation
	}
}
