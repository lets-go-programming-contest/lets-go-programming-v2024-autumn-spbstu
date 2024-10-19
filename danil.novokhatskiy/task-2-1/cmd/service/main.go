package main

import (
	"fmt"
	"log"

	"github.com/katagiriwhy/task-2-1/internal"
)

func processTemperature(k int) {
	maxTemp := 30
	minTemp := 15
	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		temp, sign, err := internal.ReadData()
		if err != nil {
			log.Fatal(err)
		}
		internal.ComputeTemp(sign, &temp, &maxTemp, &minTemp)
	}
}

func main() {
	var countDepartments int
	fmt.Print("Enter number of departments: ")
	_, err := fmt.Scan(&countDepartments)
	if countDepartments <= 0 || countDepartments > 1000 || err != nil {
		log.Fatal("Invalid number of departments")
	}
	var countWorkers int
	for i := 0; i < countDepartments; i++ {
		fmt.Print("Enter number of workers: ")
		_, err = fmt.Scan(&countWorkers)
		if countWorkers <= 0 || countWorkers > 1000 || err != nil {
			log.Fatal("Invalid number of workers")
		}
		processTemperature(countWorkers)
	}
}
