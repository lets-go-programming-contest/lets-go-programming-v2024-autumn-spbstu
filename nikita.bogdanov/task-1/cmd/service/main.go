package main

import (
	"bufio"
	"fmt"
	"os"
	"task-1/internal"
)

func main() {
	var (
		num1     string
		num2     string
		operator string
	)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите первое число: ")
	scanner.Scan()
	num1 = scanner.Text()
	//num1 = strings.ReplaceAll(num1, " ", "")
	fmt.Print("Выберите операцию (+, -, *, /): ")
	scanner.Scan()
	operator = scanner.Text()
	//operator = strings.ReplaceAll(operator, " ", "")
	fmt.Print("Введите второе число: ")
	scanner.Scan()
	num2 = scanner.Text()
	//num2 = strings.ReplaceAll(num2, " ", "")
	internal.CalcExpression(num1, num2, operator)
}
