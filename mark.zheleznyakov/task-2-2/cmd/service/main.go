package main

import (
	"container/heap"
	"fmt"

	"github.com/mrqiz/task-2-2/internal/dishesHeap"
	"github.com/mrqiz/task-2-2/internal/read"
)

func main() {
	dishes := &dishesHeap.DishesHeap{}
	heap.Init(dishes)
	n := read.ReadNumber()
	var k int
	for _ = range n {
		dish := read.ReadNumber()
		heap.Push(dishes, dish)
	}
	for {
		k = read.ReadNumber()

		if k > n {
			fmt.Println("k cannot be bigger than n")
			continue
		}

		break
	}
	dish := dishesHeap.FindDish(dishes, k)
	fmt.Println(dish)
}
