package main

import (
	"fmt"
	"log"
)

func main() {
	var op1 float64
	var op2 float64
	fmt.Println("Enter the first operand: ")
	_, err := fmt.Scanln(&op1)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	var operation string
	fmt.Println("Enter the operation: ")
	_, err = fmt.Scanln(&operation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter the second operand: ")
	_, err = fmt.Scan(&op2)
	if err != nil {
		log.Fatal(err)
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
		return
	}
	fmt.Println("Результат:", res)
}
