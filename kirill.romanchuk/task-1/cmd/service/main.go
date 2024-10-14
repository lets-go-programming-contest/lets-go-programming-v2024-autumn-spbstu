package main

import (
	"fmt"
)

func readNum(message string) float64 {
	var num float64
	for {
		fmt.Print(message)
		_, err := fmt.Scan(&num)
		if err == nil {
			break
		}
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		var dummy string
		fmt.Scanln(&dummy)
	}
	return num
}

func readOperator() string {
	var operator string
	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		_, err := fmt.Scan(&operator)
		if err == nil && isOperator(operator) {
			break
		}
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		var dummy string
		fmt.Scanln(&dummy)
	}
	return operator
}

func isOperator(operator string) bool {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	return operators[operator]
}

func calculate(num1 float64, num2 float64, operator string) (float64, error) {

	add := func(a, b float64) (float64, error) {
		return a + b, nil
	}

	subtract := func(a, b float64) (float64, error) {
		return a - b, nil
	}

	multiply := func(a, b float64) (float64, error) {
		return a * b, nil
	}

	divide := func(a, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("ошибка: деление на 0")
		}
		return a / b, nil
	}

	operations := map[string]func(float64, float64) (float64, error){
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	if opFunc, exists := operations[operator]; exists {
		return opFunc(num1, num2)
	}
	return 0, fmt.Errorf("ошибка: недопустимый оператор '%s'", operator)
}

func main() {
	for {
		fmt.Print("Введите команду (1 для вычисления, 0 для выхода): ")
		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "0":
			fmt.Println("Выход из программы.")
			return
		case "1":
			num1 := readNum("Введите первое число: ")
			operator := readOperator()
			num2 := readNum("Введите второе число: ")
			result, err := calculate(num1, num2, operator)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Результат: %.3f %s %.3f = %.3f\n", num1, operator, num2, result)
		default:
			fmt.Println("Неверная команда. Пожалуйста, попробуйте снова.")
		}
	}
}
