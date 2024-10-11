package main

import (
	"fmt"
	"os"
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

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()
	fmt.Print(readNum("test: ", 1, 2000))
}
