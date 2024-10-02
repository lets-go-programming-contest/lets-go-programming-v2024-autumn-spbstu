package main

import (
  "fmt"
  "github.com/mrqiz/task-1/internal/math"
)

func main() {
  calc := math.Calculation{
    LeftOperand: 42,
    RightOperand: 52,
    Operator: '/',
  }

  fmt.Println("Lo", calc.LeftOperand)
  fmt.Println("Ro", calc.RightOperand)
  fmt.Println("Op", string(calc.Operator))

	fmt.Printf("Res %.2f\n", math.Calculate(calc))
}

