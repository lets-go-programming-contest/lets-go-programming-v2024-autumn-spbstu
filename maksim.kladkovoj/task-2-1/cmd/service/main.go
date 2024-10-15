package main

import (
	"fmt"
	"log"

	"github.com/Mmmakskl/task-2-1/internal/calculate"
	"github.com/Mmmakskl/task-2-1/internal/input"
)

func main() {
	fmt.Print("Enter the number of departments: ")
	n, err := input.AddNumber()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < n; i++ {
		fmt.Print("Enter the number of employees: ")
		k, err := input.AddNumber()
		if err != nil {
			log.Fatal(err)
		}

		err = calculate.OptimalTemp(k)
		if err != nil {
			log.Fatal(err)
		}
	}
}
