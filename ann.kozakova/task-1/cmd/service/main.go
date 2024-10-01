package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	var a, b float32
	var operator string

	fmt.Println("enter 2 number and operation")

	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		log.Fatal("incorret number")
	}

	_, err = fmt.Scanf("%f", &b)
	if err != nil {
		log.Fatal("incorret number")
	}

	fmt.Scan(&operator)

	res, err := operate(a, b, operator)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(a, operator, b, "=", res)
}

func operate(a, b float32, operator string) (float32, error) {
	switch operator {
	case "+":
		return plus(a, b), nil
	case "-":
		return minus(a, b), nil
	case "*":
		return multiply(a, b), nil
	case "/":
		return divide(a, b)
	default:
		return 0, errors.New("no operation")
	}
}

func plus(a, b float32) float32 {
	return a + b
}

func minus(a, b float32) float32 {
	return a - b
}

func multiply(a, b float32) float32 {
	return a * b
}

func divide(a, b float32) (float32, error) {
	if b != 0 {
		return a / b, nil
	}
	return 1, errors.New("divide by 0")
}
