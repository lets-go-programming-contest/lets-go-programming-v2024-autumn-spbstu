package main

func findMax(numbers []int) int {
	max := numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{3, 7, 2, 8, 5}
	findMax(nums)
}
