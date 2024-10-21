package main

import (
	"fmt"
)

func main() {
	var topTemp uint const = 15
	var bottomTemp uint const = 30

	fmt.Println("Введите количество отделов: ")
	var departments uint
	fmt.Scan(&departments)
	
	fmt.Println("Введите количество работников отдела: ")
	var quantityOfWorkers uint
	fmt.Scan(&quantityOfWorkers)
	
	s := make([]int, 10, 15)
}