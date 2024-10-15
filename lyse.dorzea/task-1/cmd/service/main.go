package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get the first number
	fmt.Print("Введите первое число: ")
	firstInput, _ := reader.ReadString('\n')
	firstNum, err := strconv.ParseFloat(strings.TrimSpace(firstInput), 64)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	// Get the operator
	fmt.Print("Выберите операцию (+, -, *, /): ")
	operation, _ := reader.ReadString('\n')
	operation = strings.TrimSpace(operation)

	// Get the second number
	fmt.Print("Введите второе число: ")
	secondInput, _ := reader.ReadString('\n')
	secondNum, err := strconv.ParseFloat(strings.TrimSpace(secondInput), 64)
	if err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
		return
	}

	// Perform the calculation based on the operator
	switch operation {
	case "+":
		result := firstNum + secondNum
		fmt.Printf("Результат: %.2f + %.2f = %.2f\n", firstNum, secondNum, result)
	case "-":
		result := firstNum - secondNum
		fmt.Printf("Результат: %.2f - %.2f = %.2f\n", firstNum, secondNum, result)
	case "*":
		result := firstNum * secondNum
		fmt.Printf("Результат: %.2f * %.2f = %.2f\n", firstNum, secondNum, result)
	case "/":
		if secondNum == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
		} else {
			result := firstNum / secondNum
			fmt.Printf("Результат: %.2f / %.2f = %.2f\n", firstNum, secondNum, result)
		}
	default:
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}
