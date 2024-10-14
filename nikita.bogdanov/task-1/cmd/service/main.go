package main

import (
	"fmt"
	"task-1/internal"
)

func main() {
	var (
		num1     string
		num2     string
		operator string
	)
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scan(&operator)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	internal.CalcExpression(num1, num2, operator)
}
