package main

import (
	"fmt"
	"log"

	"github.com/Mmmakskl/task-8/tags/internal/calculator"
	"github.com/Mmmakskl/task-8/tags/internal/input"
)

func main() {

	var (
		number1   float64
		number2   float64
		operation string
	)

	if err := input.GetInput(&number1, &number2, &operation); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Результат: %g %s %g = %.2f\n", number1, operation, number2, calculator.Calculator(number1, number2, operation))
}
