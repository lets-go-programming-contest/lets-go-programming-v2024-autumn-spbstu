package main

import (
	"fmt"

	"github.com/Mmmakskl/task-2-1/internal/input"
	"github.com/Mmmakskl/task-2-1/internal/output"
)

func main() {
	fmt.Print("Enter the number of departments: ")
	n := input.AddNumber()

	output.OutResult(n)
}
