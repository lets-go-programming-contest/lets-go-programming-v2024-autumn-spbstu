package main

import (
	"fmt"

	"github.com/mrqiz/task-2-2/internal/heap"
)

func main() {
	dishes := &heap.DishesHeap{1, 2, 5, 6, 9, 3, 1}
	fmt.Println(heap.FindDish(dishes, 2))
}
