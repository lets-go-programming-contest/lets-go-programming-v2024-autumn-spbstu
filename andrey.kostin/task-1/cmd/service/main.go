package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	internal "github.com/IDevFrye/task-1/internal/operations"
)

var reader = bufio.NewScanner(os.Stdin)

func getOperand(prompt string) float64 {
	for {
		fmt.Print(prompt)
		reader.Scan()
		input := reader.Text()

		parts := strings.Fields(input)

		if len(parts) != 1 {
			fmt.Println("> Некорректное значение. Пожалуйста, введите одно числовое значение.")
			continue
		}

		operand, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("> Некорректное значение. Пожалуйста, введите числовое значение.")
			continue
		}

		return operand
	}
}

func getOperation() string {
	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		reader.Scan()
		input := reader.Text()

		parts := strings.Fields(input)
		if len(parts) != 1 {
			fmt.Println("> Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
			continue
		}

		operation := parts[0]
		if _, exists := internal.Operations[operation]; exists {
			return operation
		}

		fmt.Println("> Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}

func printResult(firstOperand float64, operation string, secondOperand float64, result float64) {
	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", firstOperand, operation, secondOperand, result)
}

func askContinue() bool {
	for {
		fmt.Print(">> Хотите выполнить еще одну операцию? (y/n): ")
		reader.Scan()
		input := reader.Text()
		input = strings.TrimSpace(input)

		if input == "y" {
			return true
		} else if input == "n" {
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
