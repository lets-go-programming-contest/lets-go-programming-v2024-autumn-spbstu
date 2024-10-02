package main

import (
	"fmt"
  "github.com/mrqiz/task-1/internal/input"
	"github.com/mrqiz/task-1/internal/math"
)

func main() {
	fmt.Println("i guess we doin math now")

	LeftOperand, RightOperand, Operator := input.Read()

	calculateObj := math.Calculation{
		LeftOperand,
		RightOperand,
		Operator,
	}

	result := math.Calculate(calculateObj)

	fmt.Printf("so the result is %.2f\n", result)
}

