package main

import (
	"fmt"
	"strconv"

	internal "github.com/IDevFrye/task-1/internal/operations"
)

func getOperand(prompt string) float64 {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		operand, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("> Некорректное значение. Пожалуйста, введите числовое значение.")
			continue
		}
		return operand
	}
}

func getOperation() string {
	var operation string
	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Scanln(&operation)
		if _, exists := internal.Operations[operation]; exists {
			break
		}
		fmt.Println("> Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
	return operation
}

func printResult(firstOperand float64, operation string, secondOperand float64, result float64) {
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", firstOperand, operation, secondOperand, result)
}

func askContinue() bool {
	var continueCalc string
	for {
		fmt.Print(">> Хотите выполнить еще одну операцию? (y/n): ")
		fmt.Scanln(&continueCalc)
		if continueCalc == "y" {
			return true
		} else if continueCalc == "n" {
			fmt.Println("Программа завершена.")
			return false
		} else {
			fmt.Println("> Некорректный ввод. Пожалуйста, введите 'y' или 'n'.")
		}
	}
}

func main() {
	for {
		firstNumber := getOperand("\nВведите первое число: ")
		operation := getOperation()
		secondNumber := getOperand("Введите второе число: ")

		result, err := internal.Operations[operation](firstNumber, secondNumber)
		if err != nil {
			fmt.Printf("<!> Ошибка: %s\n", err.Error())
			continue
		}
		printResult(firstNumber, operation, secondNumber, result)

		if !askContinue() {
			break
		}
	}
}
