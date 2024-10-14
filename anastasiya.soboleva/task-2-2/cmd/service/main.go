package main

import (
	"fmt"
	"log"

	"anastasiya.soboleva/task-2-2/internal/meal"
	"anastasiya.soboleva/task-2-2/internal/reader"
)

func main() {
	meals, k, err := reader.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	result := meal.FindKMeal(meals, k)

	fmt.Printf("%d\n", result)
}
