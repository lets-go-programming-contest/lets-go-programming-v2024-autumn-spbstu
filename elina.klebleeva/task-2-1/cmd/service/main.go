package main

import (
	"fmt"
	"os"
	"task-2-1/internal/calculation"
	"task-2-1/internal/input"
)

func main() {

	countOfDepart, err := input.InputCheckCount()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < countOfDepart; i++ {

		var (
			minTemp float64 = 14
			maxTemp float64 = 31
			optimum float64 = -1
		)

		countOfWorkers, err := input.InputCheckCount()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for j := 0; j < countOfWorkers; j++ {

			sign, newTemp, err := input.InputCheckNewTemp()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			err = calculation.CalculateNewTemp(&minTemp, &maxTemp, &optimum, &newTemp, sign)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if minTemp <= maxTemp {
				fmt.Println(optimum)
			} else {
				fmt.Println(-1)
				os.Exit(1)
			}

		}

	}

}
