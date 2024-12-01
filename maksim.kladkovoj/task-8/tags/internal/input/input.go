//go:build !pro

package input

import (
	"errors"
	"fmt"
)

func GetInput(number1 *float64, number2 *float64, operation *string) error {

	errOperation := errors.New("Некорректная операция. Пожалуйста, используйте символы + или -.")
	errIncNumber := errors.New("Некорректное число. Пожалуйста, введите числовое значение.")

	fmt.Println("Введите первое число: ")
	_, err := fmt.Scanln(number1)
	if err != nil {
		return fmt.Errorf("%w : %w", errIncNumber, err)
	}

	fmt.Println("Выберите операцию (+,-):")
	fmt.Scanln(operation)
	switch *operation {
	case "+", "-":
		break
	default:
		return fmt.Errorf("%w : %w", errOperation, err)
	}

	fmt.Println("Введите второе число: ")
	_, err = fmt.Scanln(number2)
	if err != nil {
		return fmt.Errorf("%w : %w", errIncNumber, err)
	}

	return nil
}
