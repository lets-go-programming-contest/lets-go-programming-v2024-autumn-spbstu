package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)-1]

	left := []int{}
	right := []int{}

	for _, el := range arr[:len(arr)-1] {
		if el <= pivot {
			left = append(left, el)
		} else {
			right = append(right, el)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}

var arr = []int{10, 7, 8, 9, 1, 5}

func main() {

	fmt.Printf("Before sort: %v\n", arr)
	sorted := quickSort(arr)
	fmt.Printf("After sort: %v\n", sorted)
}
