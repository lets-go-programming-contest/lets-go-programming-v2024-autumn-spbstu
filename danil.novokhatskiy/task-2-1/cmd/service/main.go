package main

import (
	"fmt"
	"log"
)

func processTemperature(k int) {
	var sign string
	maxTemp := 30
	minTemp := 15
	for i := 0; i < k; i++ {
		fmt.Print("Enter the temperature: ")
		fmt.Scan(&sign)
		var temp int
		_, err := fmt.Scan(&temp)
		if err != nil {
			log.Fatal(err)
		}
		//sign = strings.TrimSpace(sign)
		if sign != ">=" && sign != "<=" {
			log.Fatal("Invalid temperature")
		}
		if sign == "<=" && temp < maxTemp {
			maxTemp = temp
		}
		if sign == ">=" && temp > minTemp {
			minTemp = temp
		}
		if maxTemp >= minTemp {
			fmt.Println(minTemp)
		} else {
			fmt.Println(-1)
			break
		}
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
	for i := countDepartments; i != 0; i-- {
		fmt.Print("Enter number of workers: ")
		_, err = fmt.Scan(&countWorkers)
		if countWorkers <= 0 || countWorkers > 1000 || err != nil {
			log.Fatal("Invalid number of workers")
		}
		processTemperature(countWorkers)
	}
}
