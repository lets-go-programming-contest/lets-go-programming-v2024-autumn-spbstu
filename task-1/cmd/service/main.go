package main

import (
	"fmt"
)

func getInputData() (float64, float64, string) {
	var (
		a, b             float64
		operator, ignore string
	)

	for {
		fmt.Print("Введите первое число: ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение")
			fmt.Scanln(&ignore)
			continue
		}
		break
	}

	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Scanln(&operator)
		if operator == "+" || operator == "-" || operator == "*" || operator == "/" {
			break
		} else {
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		}
	}

	for {
		fmt.Print("Введите второе число: ")
		_, err := fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Некорректное число. Пожалуйста, введите числовое значение")
			fmt.Scanln(&ignore)
			continue
		}
		break
	}

	return a, b, operator
}

func main() {
	a, b, operator := getInputData()

	switch operator {
	case "+":
		fmt.Print("Результат: ", a+b)
	case "-":
		fmt.Print("Результат: ", a-b)
	case "*":
		fmt.Print("Результат: ", a*b)
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
		} else {
			fmt.Print("Результат: ", a/b)
		}
	}
}
