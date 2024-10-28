package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func beautifyString(str string) string {
	str = strings.TrimLeft(str, " ")
	str = strings.TrimRight(str, " ")
	return str
}

func add(num1, num2 float64) (float64, error) {
	return num1 + num2, nil
}

func subtract(num1, num2 float64) (float64, error) {
	return num1 - num2, nil
}

func multiply(num1, num2 float64) (float64, error) {
	return num1 * num2, nil
}

func divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("деление на ноль")
	}
	return num1 / num2, nil
}

type OperatorsFunc func(float64, float64) (float64, error)

func main() {
	operators := map[string]OperatorsFunc{
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}
	var input string

	for {
		fmt.Print("Введите перове число: ")
		fmt.Scanln(&input)
		num1, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка: введено не число")
			continue
		}

		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Scanln(&input)
		operator := beautifyString(input)

		fmt.Print("Введите второе число: ")
		fmt.Scanln(&input)
		num2, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка: введено не число")
			continue
		}

		if operation, exists := operators[operator]; exists {
			result, err := operation(num1, num2)
			if err != nil {
				fmt.Println("Ошибка: ", err.Error())
			} else {
				fmt.Printf("Результат: %f %s %f = %f \n", num1, operator, num2, result)
			}
		} else {
			fmt.Println("Ошибка: операция не найдена")
		}
	}

}
