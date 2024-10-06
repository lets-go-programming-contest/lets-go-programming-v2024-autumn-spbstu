package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func add(op1 float64, op2 float64) (float64, error) {
	return op1 + op2, nil
}

func minus(op1 float64, op2 float64) (float64, error) {
	return op1 - op2, nil
}

func multiply(op1 float64, op2 float64) (float64, error) {
	return op1 * op2, nil
}

func division(op1 float64, op2 float64) (float64, error) {
	if op2 == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return op1 / op2, nil
}

func readOperand(reader *bufio.Reader) (float64, error) {
	operandString, operandError := reader.ReadString('\n')
	if operandError != nil {
		return 0, operandError
	}

	operand, operandParseError := strconv.ParseFloat(strings.TrimSpace(operandString), 64)

	if operandParseError != nil {
		return 0, errors.New("некорректное число. Пожалуйста, введите числовое значение")
	}

	return operand, nil
}

func main() {
	operations := map[string]func(float64, float64) (float64, error){
		"+": add,
		"-": minus,
		"*": multiply,
		"/": division,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Введите первое число: ")

		op1, op1Error := readOperand(reader)

		if op1Error != nil {
			fmt.Println(op1Error.Error())
			continue
		}

		fmt.Print("Выберите операцию: ")

		operationSymbol, inputError := reader.ReadString('\n')

		if inputError != nil {
			fmt.Println(inputError.Error())
			continue
		}

		operationSymbol = strings.TrimSpace(strings.TrimSpace(operationSymbol))

		operation, operationExists := operations[operationSymbol]

		if !operationExists {
			fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
			continue
		}

		fmt.Print("Выберите второе число: ")

		op2, op2Error := readOperand(reader)

		if op2Error != nil {
			fmt.Println(op2Error.Error())
			continue
		}

		result, operationError := operation(op1, op2)
		if operationError != nil {
			fmt.Println("Ошибка: " + operationError.Error())
			continue
		}

		fmt.Printf("Результат: %.2f %v %.2f = %.2f\n", op1, operationSymbol, op2, result)
	}
}
