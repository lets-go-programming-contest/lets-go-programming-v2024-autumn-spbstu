package read

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/mrqiz/task-2-2/internal/dishesHeap"
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

func ReadToHeap(h *dishesHeap.DishesHeap, n int) {
	for _ = range n {
		currentDish := ReadNumber()
		heap.Push(h, currentDish)
	}
}
