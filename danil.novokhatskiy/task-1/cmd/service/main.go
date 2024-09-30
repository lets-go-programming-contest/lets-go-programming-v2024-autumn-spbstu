package main

import (
	"fmt"
)

func main() {
	var op1 float64
	var op2 float64
	fmt.Println("Enter the first operand: ")
	fmt.Scan(&op1)
	for op1 == 0 {
		fmt.Println("Please enter a number")
		fmt.Scan(&op1)
	}
	var operation string
	fmt.Println("Enter the operation: ")
	fmt.Scan(&operation)
	fmt.Println("Enter the second operand: ")
	fmt.Scan(&op2)
	for op2 == 0 {
		fmt.Println("Please enter a number")
		fmt.Scan(&op2)
	}
	var res float64
	switch operation {
	case "+":
		res = op1 + op2
	case "-":
		res = op1 - op2
	case "*":
		res = op1 * op2
	case "/":
		if op2 == 0 {
			fmt.Println("You can't divide by 0")
		} else {
			res = op1 / op2
		}
	default:
		fmt.Println("Invalid operation")
	}
	fmt.Println("Результат:", res)
}
