package input

import (
	"fmt"
	"strconv"
	"github.com/mrqiz/task-1/internal/strings"
	"github.com/mrqiz/task-1/internal/math"
)

func readNumber(label string, zeroAllowed bool) float64 {
  var input string
	var result float64

	for {
		fmt.Printf("gimme the %s number: ", label)
		fmt.Scanln(&input)

		n, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("this is not a number")
			continue
		}

		result = n

		if !zeroAllowed && result == 0 {
			fmt.Println("no zero is allowed")
			continue
		}

		break	
	}

	return result
}

func readOperator() rune {
	var operator string

	for {
		fmt.Printf("gimme an operator: ")
		fmt.Scanln(&operator)
		
		allowedOperators := []string{"+", "-", "*", "/"}

		if !strings.Has(allowedOperators, operator) {
			fmt.Println("this is not an operator, ok?")
			continue
		}

		break
	}

	return []rune(operator)[0]
}

func read() (float64, float64, rune) {
	lOperand := readNumber("first", true)
	operator := readOperator()
	rOperand := readNumber("second", string(operator) != "/")

	return lOperand, rOperand, operator
}

func ReadToCalculation(c *math.Calculation) {
	leftOperand, rightOperand, operator := read()
	c.LeftOperand = leftOperand
	c.RightOperand = rightOperand
	c.Operator = operator
}

