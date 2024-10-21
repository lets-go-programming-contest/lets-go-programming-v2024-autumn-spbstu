package reader

import (
	"container/heap"
	"fmt"
	"os"
	"task-2-2/internal/meatHeap"
)

func ReadNumber() int {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("unable to read a number")
		os.Exit(1)
	}
	return n
}

func ReadToHeap(h *meatHeap.MeatHeap, n int) {
	for range n {
		currentDish := ReadNumber()
		heap.Push(h, currentDish)
	}
}
