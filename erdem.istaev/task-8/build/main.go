package main

var arr = []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}

func main() {
	if len(arr) <= 0 {
		println("Количество чисел должно быть больше нуля.")
		return
	}

	maxValue := findMax(arr)
	println("Максимальное значение:", maxValue)
}

func findMax(arr []int) int {
	max := arr[0]
	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	return max
}
