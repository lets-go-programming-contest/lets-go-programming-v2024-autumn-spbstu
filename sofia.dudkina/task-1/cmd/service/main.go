package main

import (
	"errors"
	"fmt"
	"os"
)

func calc(n1, n2 float64, op string) (float64, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, errors.New("Division by zero")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("Invalid operation")
	}
}

func inputNumber(n *float64) {
	_, err := fmt.Scan(n)
	if err != nil {
		fmt.Println("Please enter a number")
		os.Exit(1)
	}
}

func main() {
	var n1, n2 float64
	var op string
	inputNumber(&n1)
	fmt.Scan(&op)
	inputNumber(&n2)
	res, er := calc(n1, n2, op)
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	} else {
		fmt.Println(res)
	}
}
