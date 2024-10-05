package main

import (
	"fmt"
	"log"
)

func main() {
	var op1 float64
	fmt.Print("Enter the first operand.\n> ")
	_, err := fmt.Scan(&op1)
	if err != nil {
		log.Fatal(err)
	}
	var operation string
	fmt.Println("Enter the operation: ")
	_, err = fmt.Scan(&operation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Enter the second operand: ")
	var op2 float64
	_, err = fmt.Scan(&op2)
	if err != nil {
		log.Fatal("Invalid input")
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
			log.Fatal("You can't divide by 0")
		} else {
			res = op1 / op2
		}
	default:
		log.Fatal("Invalid operation")
	}
	fmt.Println("Результат:", res)
}
