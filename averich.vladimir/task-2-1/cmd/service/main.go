package main

import (
	"fmt"
	"log"
	"math"
	"task-2-1/internal/input"
)

func main() {

	var err error

	fmt.Println("Enter the quantity of departments:")
	numberOfDepartments, err := input.ReadQuantityOfDepts()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the quantity of workers:")
	numberOfWorkers, err := input.ReadQuantityOfWorkers()
	if err != nil {
		log.Fatal(err)
	}

	rangeOfTemp := input.DiapasonOfTemperature{
		Lower: 15,
		Upper: 30,
	}

	for i := 0; i < numberOfDepartments; i++ {

		fmt.Printf("Department %v:\n", i+1)

		for k := 0; k < numberOfWorkers; k++ {

			fmt.Printf("\tWorker %v: ", k+1)

			rangeOfTemp, err = input.ReadTemperature(rangeOfTemp)
			if err != nil {
				fmt.Println("-1")
			}
		}
	}

	fmt.Println(int(math.Max(float64(rangeOfTemp.Lower), float64(rangeOfTemp.Upper))) - 1)

}
