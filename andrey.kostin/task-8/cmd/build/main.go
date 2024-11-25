package main

func BubbleSort(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func main() {
	numbers := []int{45, 23, 89, 11, 78, 4, 56, 90, 12}
	BubbleSort(numbers)
}
