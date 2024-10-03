package main

import (
	"fmt"
)

func reader() ([]float64, int) {
	var nums []float64
	var k int
	fmt.Println("Вводите числа для заполнения массива\n" +
		"(для завершения ввода нажмите Ctrl + Z (Windows) или Ctrl + D (Unix)): ")

	for {
		var num float64
		_, err := fmt.Scan(&num)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Некорректный ввод. Пожалуйста попробуйте ещё раз")
			var dummy string
			fmt.Scanln(&dummy)
		}
		nums = append(nums, num)
	}

	fmt.Println("Введите целочисленное значение k, чтобы получить k-й наибольший элемент массива:")

	for {
		_, err := fmt.Scan(&k)
		if err == nil {
			break
		}
		fmt.Println("Некорректный ввод. Пожалуйста попробуйте ещё раз")
		var dummy string
		fmt.Scanln(&dummy)
	}
	return nums, k
}

func main() {
	nums, k := reader()
	fmt.Print(nums, k)
}
