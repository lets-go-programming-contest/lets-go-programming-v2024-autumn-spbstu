package main

import (
	"fmt"
	"os"
	"task-1/internal/calculation"
	"task-1/internal/input"
)

func main() {

	a, b, operation, err := input.Input()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	result, err := calculation.Calculation(a, b, operation)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Результат : %f\n", result)

}
