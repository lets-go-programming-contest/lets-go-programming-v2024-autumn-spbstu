package main

import (
	"fmt"
	"os"
)

func maxElems(nums []int, k int) int {
	var maxNum int
	var maxIndex int

	for i := 0; i < k; i++ {
		maxNum = nums[0]
		maxIndex = 0
		for j, val := range nums {
			if val >= maxNum {
				maxNum = val
				maxIndex = j
			}
		}
		switch maxIndex {
		case len(nums) - 1:
			nums = nums[:maxIndex]
		case 0:
			nums = nums[1:]
		default:
			nums = append(nums[:maxIndex], nums[maxIndex+1:]...)
		}
	}
	return maxNum
}

func main() {
	var numLen int
	fmt.Print("Введите число элементов в массиве: ")
	_, errLen := fmt.Scan(&numLen)
	if errLen != nil || numLen <= 0 {
		fmt.Println("Некорректная длина массива")
		return
	}

	fmt.Print("Введите элементы массива через пробел: ")
	numbers := make([]int, numLen)
	for i := range numbers {
		fmt.Fscan(os.Stdin, &numbers[i])
	}

	var k int
	fmt.Print("Введите k: ")
	_, errK := fmt.Scan(&k)
	if errK != nil || k <= 0 {
		fmt.Println("Некорректное значение k")
		return
	}

	fmt.Printf("Входной массив: %v\n", numbers)
	fmt.Printf("Результат: %v", maxElems(numbers, k))
}
