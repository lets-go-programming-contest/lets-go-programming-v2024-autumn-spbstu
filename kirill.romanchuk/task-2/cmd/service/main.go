package main

import (
	"fmt"
)

func reader() ([]float64, int) {
	var nums []float64
	var k int
	fmt.Println("Вводите числа для заполнения массива\n" +
		"(для завершения ввода нажмите Ctrl + Z (Windows) или Ctrl + D (Unix)): ")

	var num float64
	for {
		_, err := fmt.Scan(&num)
		if err == nil {
			nums = append(nums, num)
			continue
		}
		if err.Error() == "EOF" {
			break
		}
		fmt.Println("Некорректный ввод. Пожалуйста попробуйте ещё раз")
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

func GetKthLargest(arr []float64, k int) float64 {
	n := len(arr)
	if k < 1 || k > n {
		panic("k должно быть в пределах длины массива")
	}
	return GetKthLargestRecursive(arr, 0, n-1, k-1)
}

func Partition(arr []float64, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func GetKthLargestRecursive(arr []float64, low, high, k int) float64 {
	if low == high {
		return arr[low]
	}
	pivotIndex := Partition(arr, low, high)

	if k == pivotIndex {
		return arr[k]
	} else if k < pivotIndex {
		return GetKthLargestRecursive(arr, low, pivotIndex-1, k)
	} else {
		return GetKthLargestRecursive(arr, pivotIndex+1, high, k)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Произошла ошибка:", r)
		}
	}()

	nums, k := reader()
	fmt.Println("input: nums = ", nums, ", k = ", k)
	result := GetKthLargest(nums, k)
	fmt.Println("Output: ", result)
}
