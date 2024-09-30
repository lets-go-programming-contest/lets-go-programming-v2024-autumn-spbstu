package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"erdem.istaev/task-1/internal/calculator"
)

var (
	ErrIncorrectNumber = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите первое число: ")
	op1, err := readFloat64(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	operation, err := readOperation(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Введите второе число: ")
	op2, err := readFloat64(reader)
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

func readFloat64(reader *bufio.Reader) (float64, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	remSpc := strings.TrimSpace(str)
	var res float64
	if res, err = strconv.ParseFloat(remSpc, 64); err != nil {
		return 0, ErrIncorrectNumber
	}
	return res, nil
}

func readOperation(reader *bufio.Reader) (string, error) {
	str, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	operation := strings.TrimSpace(str)
	if operation != "+" && operation != "-" && operation != "*" && operation != "/" {
		return "", calculator.ErrIncorrectOperation
	}
	return operation, nil
}
