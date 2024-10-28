package main

import (
	"fmt"
	"os"
	"strconv"
)

func calculator(firstNumber int, secondNumber int, oper string) (int, error) {
	switch oper {
	case "+":
		return firstNumber + secondNumber, nil
	case "-":
		return firstNumber - secondNumber, nil
	case "*":
		return firstNumber * secondNumber, nil
	case "/":
		if secondNumber == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return firstNumber / secondNumber, nil
	default:
		return 0, fmt.Errorf("неизвестный оператор")
	}
}

func main() {
	fmt.Println("Привет, это калькулятор!")

	var fistNumber int
	for {
		fmt.Print("Введи первое число: ")
		var input string
		fmt.Scan(&input)

		var err error
		fistNumber, err = strconv.Atoi(input)
		if err != nil {
			fmt.Fprint(os.Stderr, "Число введено некорректно. Пожалуйста, введите верное значение!\n")
			continue
		}
		break
	}

	var oper string
	for {
		fmt.Print("Выбери операцию (+, -, *, /): ")
		fmt.Scan(&oper)

		if oper == "" || (oper != "+" && oper != "-" && oper != "*" && oper != "/") {
			fmt.Fprint(os.Stderr, "Такая операция пока в разработке. Пожалуйста, введите доступные операции (+, -, *, /)\n")
			continue
		}
		break
	}

	var secondNumber int
	for {
		fmt.Print("Введите второе число: ")
		var input string
		fmt.Scan(&input)

		var err error
		secondNumber, err = strconv.Atoi(input)
		if err != nil {
			fmt.Fprint(os.Stderr, "Число введено некорректно. Пожалуйста, введите верное значение!\n")
			continue
		}

		break
	}

	result, err := calculator(fistNumber, secondNumber, oper)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат:", fistNumber, oper, secondNumber, "=", result)
	}
}
