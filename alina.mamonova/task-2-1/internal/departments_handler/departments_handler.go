package departments_handler

import (
	"bufio"
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

func Start(reader *bufio.Reader) {
	departmentsNum := input.ProcessNumInput("Enter the number of departments: ", reader)
	readDepartment(departmentsNum, reader)

}

func readDepartment(departmentsNum int, reader *bufio.Reader) {
	for i := 0; i < departmentsNum; i++ {
		fmt.Printf("\nProcessing department %d:\n", i+1)
		employeesNum := input.ProcessNumInput("Enter the number of employees: ", reader)
		processDepartment(employeesNum, reader)
	}
}

func processDepartment(employeesNum int, reader *bufio.Reader) {
	lowerBound := minTemp
	upperBound := maxTemp
	readEmployees(employeesNum, lowerBound, upperBound, reader)
}

func readEmployees(employeesNum int, lowerBound int, upperBound int, reader *bufio.Reader) {
	for j := 0; j < employeesNum; j++ {
		operator, temp := input.ProcessFullInput(reader)
		processEmployees(operator, temp, lowerBound, upperBound)
	}
}

func processEmployees(operator string, temp int, lowerBound int, upperBound int) {
	updateOptimalTemp(operator, temp, &lowerBound, &upperBound)
	if lowerBound <= upperBound {
		fmt.Printf("Optimal temperature for department: %d°C\n", lowerBound)
	} else {
		fmt.Println(-1)
		os.Exit(1)

	}
}
