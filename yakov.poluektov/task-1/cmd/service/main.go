package main

import (
	"task-1/internal/calculator"
)

func main() {
	calc := calculator.NewConsoleCalculator()
	calc.Run()
}
