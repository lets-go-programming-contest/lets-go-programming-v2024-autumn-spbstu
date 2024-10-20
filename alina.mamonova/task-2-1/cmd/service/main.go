package main

import (
	"fmt"
	"github.com/hahapathetic/task-2-1/internal/input"
	"os"
)

const (
	minTemp = 15
	maxTemp = 30
)

func updateOptimalTemp(operator string, curTemp int, lowerBound, upperBound *int) {
	switch operator {
	case "<=":
		if curTemp < *upperBound {
			*upperBound = curTemp
		}
	case ">=":
		if curTemp > *lowerBound {
			*lowerBound = curTemp
		}
	}
}

func calculateOptimalTemp() {
	departmentsNum := input.ProcessNumInput("Enter the number of departments: ")

	for i := 0; i < departmentsNum; i++ {
		fmt.Printf("\nProcessing department %d:\n", i+1)
		employeesNum := input.ProcessNumInput("Enter the number of employees: ")
		lowerBound := minTemp
		upperBound := maxTemp

		for j := 0; j < employeesNum; j++ {
			operator, temp := input.ProcessFullInput()
			updateOptimalTemp(operator, temp, &lowerBound, &upperBound)
			if lowerBound <= upperBound {
				fmt.Printf("Optimal temperature for department %d: %d°C\n", i+1, lowerBound)
			} else {
				fmt.Println(-1)
				os.Exit(1)

			}
		}
	}
}

func main() {
	fmt.Println("Welcome to Optimal temperature handler!")
	calculateOptimalTemp()
}
