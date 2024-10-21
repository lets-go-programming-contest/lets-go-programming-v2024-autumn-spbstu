package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/katagiriwhy/task-2-1/internal"
)

func processTemperature(k int) (err error) {
	maxTemp := 30
	minTemp := 15
	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		temp, sign, err := internal.ReadData()
		if err != nil {
			return err
		}
		err = internal.ComputeTemp(sign, &temp, &maxTemp, &minTemp)
		if err != nil {
			return err
		}
	}
	return nil
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
		err = processTemperature(countWorkers)
		if errors.Is(err, internal.ErrorTemp{}) {
			fmt.Println("-1")
			os.Exit(1)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
