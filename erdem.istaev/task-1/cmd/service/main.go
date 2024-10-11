package main

import (
	"bufio"
	"fmt"
	"os"

	"erdem.istaev/task-1/internal/calculator"
	"erdem.istaev/task-1/internal/input"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите первое число: ")
	op1, err := input.ReadFloat64(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	operation, err := input.ReadOperation(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите второе число: ")
	op2, err := input.ReadFloat64(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := calculator.Calculate(op1, op2, operation)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Результат: %g %s %g = %g", op1, operation, op2, res)
}
