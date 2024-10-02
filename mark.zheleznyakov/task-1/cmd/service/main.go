package main

import (
	"fmt"
  "github.com/mrqiz/task-1/internal/input"
	"github.com/mrqiz/task-1/internal/math"
)

func main() {
	fmt.Println("i guess we doin math now")
	calculation := math.Calculation{}
	input.ReadToCalculation(&calculation)
	result := math.Calculate(calculation)
	fmt.Printf("so the result is %.2f\n", result)
}

