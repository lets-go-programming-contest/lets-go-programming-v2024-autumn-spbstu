package main

import "fmt"

func main() {
	var n, k int
	fmt.Print("Enter the number of elements in the array: ")
	fmt.Scan(&n)
	nums := make([]int, n)
	fmt.Print("Enter the elements: ")
	for i := range n {
		fmt.Scan(&nums[i])
	}
	fmt.Print("Enter k: ")
	fmt.Scan(&k)
	QuickSort(nums)
	fmt.Printf("Sorted array: %v\n",nums)
	fmt.Printf("The largest k-th element in an array: %d\n", nums[len(nums)-k])
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, low int, high int) {
	if low < high {
		pi := partition(nums, low, high)
		quickSort(nums, low, pi-1)
		quickSort(nums, pi+1, high)
	}
}

func partition(nums []int, low int, high int) int {
	pivot := nums[high]
	i := low - 1
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}
