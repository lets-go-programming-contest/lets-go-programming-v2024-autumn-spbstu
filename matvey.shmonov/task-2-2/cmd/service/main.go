package main

import (
	"container/heap"
	"fmt"
	"log"

	"github.com/Koshsky/task-2-2/internal/structure"
)

func main() {
	n := ReadInt("Enter the number of dishes: ")
	var nums structure.IntHeap
	fmt.Print("Enter the elements: ")
	for i := range n {
		fmt.Scan(&i)
		heap.Push(&nums, i)
	}
	k := ReadInt("Enter k: ")
	fmt.Printf("The best k-th dish in an array: %d\n", kMax(nums, k))
}

func ReadInt(message string) int {
	fmt.Print(message)
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return input
}

func kMax(rating structure.IntHeap, k int) int {
	for i := 1; i < k; i++ {
		heap.Pop(&rating)
	}
	return rating[0]
}
