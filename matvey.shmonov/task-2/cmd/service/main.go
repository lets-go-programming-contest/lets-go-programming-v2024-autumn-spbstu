package main

import "fmt"

const ARRAY_SIZE = 10

func main() {
	nums := [ARRAY_SIZE]int{}
	var k int
	for i := range ARRAY_SIZE {
		fmt.Scan(&nums[i])
	}
	fmt.Scan(&k)
	QuickSort(nums[:])
	fmt.Println(nums)
	fmt.Println(nums[len(nums) - k])
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
