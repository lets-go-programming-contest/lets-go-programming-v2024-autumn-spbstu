package main

import (
	"fmt"
)

func getInputData() (float64, float64, string) {
	var (
		a, b              float64
		operation, ignore string
	)

	for {
		fmt.Print("Введите первое число: ")
		_, err := fmt.Scan(&a)
		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
			fmt.Scan(&ignore)
			continue
		}
		break
	}

	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Scan(&operation)
		if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		} else {
			break
		}
	}

	for {
		fmt.Print("Введите второе число: ")
		_, err := fmt.Scan(&b)
		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
			fmt.Scanln(&ignore)
			continue
		}
		break
	}

	return a, b, operation
}

func main() {
	a, b, operation := getInputData()

	switch operation {
	case "+":
		fmt.Println("Результат: ", a+b)
	case "-":
		fmt.Println("Результат: ", a-b)
	case "*":
		fmt.Println("Результат: ", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
		} else {
			fmt.Println("Результат:  ", a/b)
		}
	}
}
