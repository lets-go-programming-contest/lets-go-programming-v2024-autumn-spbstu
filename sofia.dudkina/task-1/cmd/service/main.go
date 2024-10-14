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
			return 0, errors.New("division by zero")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("invalid operation")
	}
}

func inputNumber(n *float64) error {
	_, err := fmt.Scanln(n)
	if err != nil {
		return errors.New("incorrect input")
	} else {
		return nil
	}
}

func main() {
	var n1, n2 float64
	var op string
	fmt.Println("Enter number: ")
	err := inputNumber(&n1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Enter operation: ")
	_, err = fmt.Scanln(&op)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Enter number: ")
	err = inputNumber(&n2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, er := calc(n1, n2, op)
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	} else {
		fmt.Println(res)
	}
}
