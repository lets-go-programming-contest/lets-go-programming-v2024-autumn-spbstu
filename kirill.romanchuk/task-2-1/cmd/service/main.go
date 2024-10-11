package main

import (
	"fmt"
	"os"
)

const (
	LowerBound = 15
	UpperBound = 30
)

func readNum(message string, min int, max int) int {
	var num int
	fmt.Print(message)
	_, err := fmt.Scan(&num)
	if err != nil {
		panic("ошибка: некорректное значение")
	}
	if num < min || num > max {
		panic(fmt.Sprintf("ошибка: значение должно быть в диапазоне от %d до %d", min, max))
	}
	return num
}

func readConditionAndTemperature() (string, int) {
	var condition string
	fmt.Print("Введите условие (>= или <=) и затем температуру (15-30): ")
	_, err := fmt.Scan(&condition)
	if err != nil {
		panic("ошибка: некорректное значение")
	}
	if condition != "<=" && condition != ">=" { //map?
		panic("Неверное условие: должно быть '<=' или '>='")
	}
	return condition, readNum("", LowerBound, UpperBound)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	fmt.Print(readNum("test: ", 1, 2000))
	fmt.Println()
	fmt.Print(readConditionAndTemperature())
}
