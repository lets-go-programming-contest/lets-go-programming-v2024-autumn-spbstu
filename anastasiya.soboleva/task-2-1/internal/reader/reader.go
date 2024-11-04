package reader

import (
	"fmt"
	"log"
)

func ReadDepartmentCount() int {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil || n < 1 || n > 1000 {
		log.Fatal("Ошибка: некорректное количество отделов")
	}
	return n
}

func ReadEmployeeCount() int {
	var k int
	_, err := fmt.Scanln(&k)
	if err != nil || k < 1 || k > 1000 {
		log.Fatal("Ошибка: некорректное количество сотрудников в отделе")
	}
	return k
}

func ReadTemperatureOperation() (string, int) {
	var operation string
	var value int
	_, err := fmt.Scanf("%s %d\n", &operation, &value)
	if err != nil || value > 30 || value < 15 || (operation != ">=" && operation != "<=") {
		log.Fatal("Ошибка: некорректный ввод данных")
	}
	return operation, value
}
