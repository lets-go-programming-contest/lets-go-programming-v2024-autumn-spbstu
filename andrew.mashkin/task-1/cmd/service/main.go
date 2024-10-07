package main

import (
	"fmt"
	"log"
)

func main() {
	var first int
	var second int
	var operation string
	_, err := fmt.Scan(&first, &second)
	if err != nil {
		log.Fatal("Некорректное число. Пожалуйста, введите числовое значение.")
	}
	_, err = fmt.Scan(&operation)
	if err != nil {
		log.Fatal(err)
	}
	result, err := Calcualte(first, second, operation)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
