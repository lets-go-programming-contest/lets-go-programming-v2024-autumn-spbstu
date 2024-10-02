package main

import (
	"fmt"
	"github.com/mrqiz/task-1/internal/structures"
)

func main() {	
	calc := structures.Calculation{
		LeftOperand: 42,
		RightOperand: 52,
		Operator: '+',
	}

	fmt.Println("Lo", calc.LeftOperand)
	fmt.Println("Ro", calc.RightOperand)
	fmt.Println("Op", string(calc.Operator))
}

