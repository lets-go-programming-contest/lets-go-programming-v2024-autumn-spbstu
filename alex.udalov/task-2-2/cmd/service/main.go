package main

import (
	"container/heap"
	"fmt"
	"task-2-2/internal/meatHeap"
	"task-2-2/internal/reader"
)

func main() {
	dishes := &meatHeap.MeatHeap{}
	heap.Init(dishes) //инициализация кучи
	n := reader.ReadNumber()
	reader.ReadToHeap(dishes, n)
	var k int
	for {
		k = reader.ReadNumber()
		if k > n {
			fmt.Println("k не может быть больше n")
			continue
		}
		break
	}
	dish := meatHeap.FindMeat(dishes, k)
	fmt.Println(dish)
}
