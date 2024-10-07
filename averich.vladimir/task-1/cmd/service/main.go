package main

import(
	"fmt"
	"strconv"
	"errors"
)

func readOperand(input string) (float64, error) {

	var number float64
	number, err := strconv.ParseFloat(input, 64)

	if err == nil {
		return number, err
	} else {
		return 0, errors.New("введите число")
	}
}

func evaluateExpression(firstOperand float64, secondOperand float64, operator string) (float64, error) {
	var result float64
	var err error

	switch operator {
		case "+":
			result = firstOperand + secondOperand
		case "-":
			result = firstOperand - secondOperand
		case "*":
			result = firstOperand * secondOperand
		case "/":
			if(secondOperand != 0) {
				result = firstOperand / secondOperand
			} else {
				err = errors.New("делить на 0 нельзя")
				return result, err
			}
		default:
			err = errors.New("введите оператор")
	}
	return result, err
}

func main() {
	var err error
	var input, operator string
	var firstOperand, secondOperand, result float64

	for{

		if input == "хватит" {
			fmt.Println("До новых встреч :)")
			return
		}

		fmt.Print("Введите первое число: ")
		fmt.Scan(&input)
		firstOperand, err = readOperand(input)

		if err != nil {
			fmt.Println("Некорректный ввод:", err)
			continue
		}

		fmt.Print("Введите оператор: ")
		fmt.Scan(&input)
		operator = input

		fmt.Print("Введите второе число: ")
		fmt.Scan(&input)
		secondOperand, err = readOperand(input)

		if err != nil {
			fmt.Println("Некорректный ввод:", err)
			continue
		}

		result, err = evaluateExpression(firstOperand, secondOperand, operator)

		if err != nil {
			fmt.Println("Некорректный ввод:", err)
			continue
		}

		fmt.Println("Результат: ", result)
	}
}
