package main

import (
	"fmt"
)

func main() {
	const topTemp uint = 15
	const bottomTemp uint = 30

	fmt.Println("Введите количество отделов: ")
	var departments int
	fmt.Scan(&departments)

	for i := 0; i < departments; i++ {
		fmt.Println("Введите количество работников отдела: ")
		var quantityOfWorkers uint
		fmt.Scan(&quantityOfWorkers)
	}
	
}