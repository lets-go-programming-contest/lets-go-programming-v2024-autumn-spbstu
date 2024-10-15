package main

import (
	"fmt"

	"github.com/Madyarov-Gleb/task-2-1/internal/input"
	"github.com/Madyarov-Gleb/task-2-1/internal/output"
)

func main() {
	var (
		dep       int
		emp       int
		tempSign  string
		tempValue int
		maxHigher int
		minLower  int
	)
	dep = input.AddNumber()
	for i := 1; i <= dep; i++ {
		emp = input.AddNumber()
		maxHigher = 0
		minLower = 100
		for j := 1; j <= emp; j++ {
			tempSign = input.AddTempSign()
			tempValue = input.AddNumber()
			if tempValue < 15 || tempValue > 30 {
				fmt.Println("Invalid temperature value, expected value between 15 and 30")
				break
			}
			if tempSign == "<=" && tempValue < minLower {
				minLower = tempValue
			}
			if tempSign == ">=" && tempValue > maxHigher {
				maxHigher = tempValue
			}
			fmt.Println(output.OutAnswer(maxHigher, minLower))
		}
	}
}
