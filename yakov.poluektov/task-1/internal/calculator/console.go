package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		return 0, err
	}

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func readAnswerFromConsole() (bool, error) {
	in := bufio.NewReader(os.Stdin)
	rowsInput, err := in.ReadString('\n')
	if err != nil {
		return false, err
	}

	input := strings.TrimSpace(rowsInput)

	if input == "Yes" {
		return true, nil
	} else if input == "No" {
		return false, nil
	}
	return false, fmt.Errorf("incorrect responce")
}

func getOperationBySymbolFromConsole() (func(o1, o2 float64) (float64, error), error) {
	in := bufio.NewReader(os.Stdin)
	num, err := in.ReadString('\n')
	if err != nil {
		return nil, err
	}

	switch strings.TrimSpace(num) {
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
