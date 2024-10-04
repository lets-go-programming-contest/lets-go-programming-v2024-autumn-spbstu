package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func findKElement(arr []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return heap.Pop(h).(int)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите числа, разделенные пробелами:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	numbers := strings.Split(input, " ")

	var arr []int
	for _, numberStr := range numbers {
		number, _ := strconv.Atoi(numberStr)
		arr = append(arr, number)
	}
	fmt.Println("Введите k: ")
	fmt.Print("> ")
	var k int
	_, err := fmt.Fscanln(reader, &k)
	if err != nil {
		log.Fatal(err)
		return
	}
	if k > 0 && k <= len(arr) {
		result := findKElement(arr, k)
		fmt.Printf("%d-й наибольший элемент массива: %d\n", k, result)
	} else {
		fmt.Println("Некорректное значение k.")
	}
}
