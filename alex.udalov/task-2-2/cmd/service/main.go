package main

import (
	"container/heap"
	"fmt"
	"os"

	"task-2-2/internal/meatHeap"
	"task-2-2/internal/reader"
)

func main() {
	dishes := &meatHeap.MeatHeap{}
	heap.Init(dishes)
	n := reader.ReadNumber()
	reader.ReadToHeap(dishes, n)
	k := reader.ReadNumber()
	if k > n {
		fmt.Println("k не может быть больше n")
		os.Exit(1)
	}
	dish := meatHeap.FindMeat(dishes, k)
	fmt.Println(dish)
}
