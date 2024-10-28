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

func start() {
	departmentsNum := input.ProcessNumInput("Enter the number of departments: ")
	readDep(departmentsNum)

}

func readDep(departmentsNum int) {
	for i := 0; i < departmentsNum; i++ {
		fmt.Printf("\nProcessing department %d:\n", i+1)
		employeesNum := input.ProcessNumInput("Enter the number of employees: ")
		processDep(employeesNum)
	}
}

func processDep(employeesNum int) {
	lowerBound := minTemp
	upperBound := maxTemp
	readEmp(employeesNum, lowerBound, upperBound)
}

func readEmp(employeesNum int, lowerBound int, upperBound int) {
	for j := 0; j < employeesNum; j++ {
		operator, temp := input.ProcessFullInput()
		processEmp(operator, temp, lowerBound, upperBound)
	}
}

func processEmp(operator string, temp int, lowerBound int, upperBound int) {
	updateOptimalTemp(operator, temp, &lowerBound, &upperBound)
	if lowerBound <= upperBound {
		fmt.Printf("Optimal temperature for department: %d°C\n", lowerBound)
	} else {
		fmt.Println(-1)
		os.Exit(1)

	}
}

func main() {
	fmt.Println("Welcome to Optimal temperature handler!")
	start()
}
