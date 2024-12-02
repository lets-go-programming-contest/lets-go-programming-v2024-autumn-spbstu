//go:build pro
// +build pro

package main

import "fmt"

func findMin(numbers []int) int {
	min := numbers[0]
	for _, num := range numbers {
		if num < min {
			min = num
		}
	}
	return min
}

func init() {
	nums := []int{3, 7, 2, 8, 5}
	fmt.Println("Минимальное число:", findMin(nums))
}
