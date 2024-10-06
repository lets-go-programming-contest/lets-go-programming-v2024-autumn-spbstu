package calculator

import (
	"fmt"
	"strconv"
)

type ConsoleCalculator struct {
}

func NewConsoleCalculator() *ConsoleCalculator {
	return &ConsoleCalculator{}
}

func (c *ConsoleCalculator) Run() {
	for {
		fmt.Println("Enter first number:")
		first, err := readNumberFromConsole()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		fmt.Println("Enter operation:")
		op, err := getOperationBySymbolFromConsole()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		fmt.Println("Enter second number:")
		second, err := readNumberFromConsole()
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		res, err := op(first, second)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}

		fmt.Println("Result:", res)

		fmt.Println("Wanna continue?(Yes/No):")
		cont, err := readAnswerFromConsole()
		if err != nil {
			fmt.Println("Error:", err.Error())
			break
		}
		if !cont {
			break
		}
	}
}

func readNumberFromConsole() (float64, error) {
	var numberString string
	_, err := fmt.Scan(&numberString)
	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseFloat(numberString, 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func readAnswerFromConsole() (bool, error) {
	var answer string
	if _, err := fmt.Scan(&answer); err != nil {
		return false, err
	}

	if answer == "Yes" {
		return true, nil
	} else if answer == "No" {
		return false, nil
	}
	return false, fmt.Errorf("incorrect responce")
}

func getOperationBySymbolFromConsole() (func(o1, o2 float64) (float64, error), error) {
	var num string
	_, err := fmt.Scan(&num)
	if err != nil {
		return nil, err
	}

	switch num {
	case "-":
		return func(o1, o2 float64) (float64, error) {
			return o1 - o2, nil
		}, nil
	case "+":
		return func(o1, o2 float64) (float64, error) {
			return o1 + o2, nil
		}, nil
	case "*":
		return func(o1, o2 float64) (float64, error) {
			return o1 * o2, nil
		}, nil
	case "/":
		return func(o1, o2 float64) (float64, error) {
			if o2 == 0 {
				return 0, fmt.Errorf("cannot divide by zero")
			}
			return o1 / o2, nil
		}, nil
	}

	return nil, fmt.Errorf("unknown operation")
}
