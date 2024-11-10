package main

import (
	"container/heap"
	"fmt"
	"os"

	"github.com/kirill.romanchuk/task-2-2/internal/maxheap"
	"github.com/kirill.romanchuk/task-2-2/internal/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
			os.Exit(1)
		}
	}()

	_, values, k := utils.ReadInput()

	var nums maxheap.IntMaxHeap

	for _, value := range values {
		heap.Push(&nums, value)
	}

	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(&nums).(int)
	}

	fmt.Println(result)
}
