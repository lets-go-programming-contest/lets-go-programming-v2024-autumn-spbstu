package utils

import (
	"fmt"
)

func ReadIntNum(message string, min int, max int) int {
	var num float64
	if message != "" {
		fmt.Print(message)
	}
	_, err := fmt.Scan(&num)
	if err != nil {
		panic("ошибка: некорректное значение")
	}
	resultNum := int(num)
	if num != float64(resultNum) {
		panic("ошибка: число должно быть целым")
	}
	if resultNum < min || resultNum > max {
		panic(fmt.Sprintf("ошибка: значение должно быть в диапазоне от %d до %d", min, max))
	}
	return resultNum
}

func ReadInput() (int, []int, int) {
	n := ReadIntNum("Введите количество блюд (от 1 до 10000): ", 1, 10000)
	values := make([]int, n)
	fmt.Printf("Введите значение для блюд (от -10000 до 10000): ")

	for i := 0; i < n; i++ {
		values[i] = ReadIntNum("", -10000, 10000)
	}

	k := ReadIntNum(fmt.Sprintf("Введите номер предпочитаемого блюда (от 1 до %d): ", n), 1, n)

	return n, values, k
}
