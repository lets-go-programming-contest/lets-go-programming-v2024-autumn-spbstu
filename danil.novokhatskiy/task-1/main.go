package main

import (
	"fmt"
	"strconv"
)

func main() {
	var op1 int
	var op2 int
	var operation string
	fmt.Println("Enter the first operand: ")
	fmt.Scan(&op1)
	_, err := strconv.Atoi(strconv.Itoa(op1))
	for err == nil {
		fmt.Println("Please enter a number")
		fmt.Scan(&op1)
		_, err = strconv.Atoi(strconv.Itoa(op1))
	}
	fmt.Println("Enter the operation: ")
	fmt.Scan(&operation)
	fmt.Println("Enter the second operand: ")
	fmt.Scan(&op2)
	_, err = strconv.Atoi(strconv.Itoa(op2))
	for err == nil {
		fmt.Println("Please enter a number")
		fmt.Scan(&op2)
		_, err = strconv.Atoi(strconv.Itoa(op2))
	}
	var res int
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
	case "%":
		res = op1 % op2
	default:
		fmt.Println("Invalid operation")
	}
	fmt.Println("Результат:", res)
}
