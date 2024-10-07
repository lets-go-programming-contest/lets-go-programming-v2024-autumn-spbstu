package input

import (
	"errors"
	"fmt"
	"strconv"
)

func inputNumber(prompt string) (float64, error) {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, errors.New("некорректный ввод операндов, пожалуйста, используйте числа")
	}
	return number, nil
}

func Input() (float64, float64, string, error) {
	fmt.Println("Приветствуем в программе 'Калькулятор'")

	x, err := inputNumber("Введите первое число: ")
	if err != nil {
		return 0, 0, "", err
	}

	var operation string
	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scan(&operation)

	y, err := inputNumber("Введите второе число: ")
	if err != nil {
		return 0, 0, "", err
	}

	return x, y, operation, nil
}
