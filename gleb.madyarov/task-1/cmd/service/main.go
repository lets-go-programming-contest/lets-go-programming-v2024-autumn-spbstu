package main

import (
	"fmt"
)

func main() {
	var op1 float64
	var op2 float64
	var operation string
	var answer float64
	var choice int
	var input string
	for {
		fmt.Print("enter 1 to calculate or 0 to exit: ")
		_, err1 := fmt.Scanln(&choice)
		if err1 != nil {
			fmt.Println("Invalid input")
			fmt.Scanln(&input)
			continue
		}
		if choice == 0 {
			break
		}
		for {
			fmt.Print("enter the first operand: ")
			_, err2 := fmt.Scan(&op1)
			if err2 != nil {
				fmt.Scanln(&input)
				fmt.Println("Invalid input, please try again")
				continue
			}
			break
		}
		for {
			fmt.Print("enter the operation: ")
			fmt.Scan(&input)
			switch input {
			case "+":
				operation = input
			case "-":
				operation = input
			case "*":
				operation = input
			case "/":
				operation = input
			default:
				fmt.Println("Invalid input, please try again")
				continue
			}
			break
		}
		for {
			fmt.Print("enter the second operand: ")
			_, err3 := fmt.Scan(&op2)
			if err3 != nil {
				fmt.Scanln(&input)
				fmt.Println("Invalid input, please try again")
				continue
			}
			break
		}
		switch operation {
		case "+":
			answer = op1 + op2
			fmt.Println(answer)
		case "-":
			answer = op1 - op2
			fmt.Println(answer)
		case "*":
			answer = op1 * op2
			fmt.Println(answer)
		case "/":
			if op2 == 0 {
				fmt.Println("Division by 0 is impossible")
				break
			}
			answer = op1 / op2
			fmt.Println(answer)
		default:
			fmt.Println("Error")
		}
	}
}
