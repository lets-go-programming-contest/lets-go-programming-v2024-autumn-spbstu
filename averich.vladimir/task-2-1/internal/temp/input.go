package input

import (
	"fmt"
	"log"
)

func safeInputQuantityOfDepts() int {
	var number int

	_, err := fmt.Scanln(&number)
	if err != nil || number < 1 || number > 1000 {
		log.Fatal("Слишком много отделов, либо некорректный ввод")
	}

	return number
}

func safeInputQuantityOfWorkers() int {
	var number int

	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 || n > 1000 {
		log.Fatal("Слишком много сотрудников, либо некорректный ввод")
	}

	return number
}

func safeInputOfTemperature() (string, int) {
	var temperature int
	var operation string
	
	_, err := fmt.Scanf("%s %d\n", &operation, &temperature)
	if err != nil || k > 30 || k < 15 || (operator != ">=" && operator != "<=") {
		log.Fatal("Некорректный ввод")
	}

	return operation, temperature
}