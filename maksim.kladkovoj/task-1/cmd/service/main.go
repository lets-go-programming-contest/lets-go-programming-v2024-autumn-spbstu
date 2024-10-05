package main

import (
	"errors"
	"fmt"
	"log"
)

func getInput() (float64, float64, string) {
	var (
		number1   float64
		number2   float64
		operation string
	)

	errOperation := errors.New("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	errDivZero := errors.New("Ошибка: деление на ноль невозможно.")
	errIncNumber := errors.New("Некорректное число. Пожалуйста, введите числовое значение.")

	fmt.Println("Введите первое число: ")
	_, err := fmt.Scanln(&number1)
	if err != nil {
		log.Fatal(errIncNumber)
	}

	fmt.Println("Выберите операцию (+,-,*,/):")
	fmt.Scanln(&operation)
	switch operation {
	case "+", "-", "*", "/":
		break
	default:
		log.Fatal(errOperation)
	}

	fmt.Println("Введите второе число: ")
	_, err = fmt.Scanln(&number2)
	if err != nil {
		log.Fatal(errIncNumber)
	} else if operation == "/" && number2 == 0 {
		log.Fatal(errDivZero)
	}

	return number1, number2, operation
}

func calculator(number1 float64, number2 float64, operation string) float64 {
	switch operation {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "*":
		return number1 * number2
	default:
		return number1 / number2
	}
}

func main() {

	number1, number2, operation := getInput()

	fmt.Printf("%s %g %s %g %s %.2f %s", "Результат:", number1, operation, number2, "=", calculator(number1, number2, operation), "\n")
}
