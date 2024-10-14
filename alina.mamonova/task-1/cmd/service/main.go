package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func addition(a float64, b float64) float64 {
	return a + b
}

func subtraction(a float64, b float64) float64 {
	return a - b
}

func multiplication(a float64, b float64) float64 {
	return a * b
}

func division(a float64, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		os.Exit(1)
	}
	return a / b
}

func calculate(a float64, b float64, operator string) float64 {
	var result float64 = 0

	switch operator {
	case "+":
		result = addition(a, b)
	case "-":
		result = subtraction(a, b)
	case "*":
		result = multiplication(a, b)
	case "/":
		result = division(a, b)
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}
	return result
}

func processOperandInput(intro string) float64 {
	var operand float64 = 0
	for {
		fmt.Print(intro)

		reader := bufio.NewReader(os.Stdin)
		input, err1 := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err1 == nil {
			num, err2 := strconv.ParseFloat(input, 64)
			if err2 == nil {
				operand = num
				break
			} else {
				fmt.Println("Invalid input. Please enter a valid number.")
			}
		} else {
			fmt.Println("Invalid input. Please enter a valid number.")
		}
	}
	return operand
}

func processOperatorInput(intro string) string {
	var exp = ""
	fmt.Print(intro)
	for {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if err == nil {

			if input == "+" || input == "-" || input == "*" || input == "/" {
				exp = input
				break
			}
		}
		fmt.Println("Invalid operator. Please enter one of the following: +, -, *, /.")
	}
	return exp
}

func main() {

	var firstOperand float64 = 0
	var secondOperand float64 = 0
	var operator = ""

	fmt.Println("Welcome to calculator!")

	firstOperand = processOperandInput("Enter the first number: ")
	operator = processOperatorInput("Enter the operator: ")
	secondOperand = processOperandInput("Enter the second number: ")

	fmt.Println("Result:", firstOperand, operator, secondOperand, "=", calculate(firstOperand, secondOperand, operator))
}
