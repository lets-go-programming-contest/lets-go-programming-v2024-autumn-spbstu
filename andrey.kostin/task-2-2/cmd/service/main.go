package main

import (
	"fmt"
	"log"

	"github.com/IDevFrye/task-2-2/internal/findDish"
	reader "github.com/IDevFrye/task-2-2/internal/io"
)

var (
	minN_k = 1
	maxN   = 10000
	minAi  = -10000
	maxAi  = 10000
)

func main() {
	countOfDishes, err := reader.GetNumber("Введите количество блюд на шведском столе: ", minN_k, maxN)
	if err != nil {
		log.Fatal(err)
	}
	preferences, err := reader.GetHeapElements("Введите предпочтение к блюду: ", minAi, maxAi, countOfDishes)
	if err != nil {
		log.Fatal(err)
	}
	numberOfPreference, err := reader.GetPrefDish("Введите порядковый номер блюда по предпочтению: ", minN_k, countOfDishes)
	if err != nil {
		log.Fatal(err)
	}
	result := findDish.FindDish(&preferences, numberOfPreference)
	fmt.Println("Число, соответствующее", numberOfPreference, "- му наиболее предпочтительному блюду:", result)
}
