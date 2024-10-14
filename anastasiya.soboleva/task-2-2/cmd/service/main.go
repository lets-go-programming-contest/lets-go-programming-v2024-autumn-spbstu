package main

import (
	"anastasiya.soboleva/task-2-2/internal/meal"
	"anastasiya.soboleva/task-2-2/internal/reader"
	"fmt"
	"log"
)

func main() {
	meals, k, err := reader.ReadInput()
	if err != nil {
		log.Fatal(err)
	}

	result := meal.FindKMeal(meals, k)

	fmt.Printf("%d\n", result)
}
