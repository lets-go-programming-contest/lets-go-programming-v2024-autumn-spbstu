package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/katagiriwhy/task-2-1/internal"
)

const (
	maxWorkersAndDep = 1000
	minWorkersAndDep = 0
)

func main() {
	var countDepartments int
	fmt.Print("Enter number of departments: ")
	_, err := fmt.Scan(&countDepartments)
	if countDepartments <= minWorkersAndDep || countDepartments > maxWorkersAndDep || err != nil {
		log.Fatal("Invalid number of departments")
	}
	var countWorkers int
	for i := 0; i < countDepartments; i++ {
		fmt.Print("Enter number of workers: ")
		_, err = fmt.Scan(&countWorkers)
		if countWorkers <= minWorkersAndDep || countWorkers > maxWorkersAndDep || err != nil {
			log.Fatal("Invalid number of workers")
		}
		maxTemp := 30
		minTemp := 15
		err = internal.GetTemp(countWorkers, os.Stdout, &maxTemp, &minTemp)
		if errors.Is(err, internal.ErrorTemp{}) {
			fmt.Println("-1")
			os.Exit(1)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
