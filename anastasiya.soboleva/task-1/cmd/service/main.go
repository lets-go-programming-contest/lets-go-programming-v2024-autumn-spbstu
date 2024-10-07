package main

import (
	"errors"
	"fmt"
	"log"
)

func Operation(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("делить на ноль нельзя")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("некорректная операция: %s", operator)
	}
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		log.Fatalf("Ошибка: некорректное первое число. %v", err)
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		log.Fatalf("Ошибка: некорректное второе число. %v", err)
	}

	fmt.Print("Введите операцию (+, -, *, /): ")
	_, err = fmt.Scan(&operator)

	result, err := Operation(num1, num2, operator)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	fmt.Printf("Результат: %.2f\n", result)
}
