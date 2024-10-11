package readDishes

import (
	"container/heap"
	"fmt"

	"github.com/artem6554/task-2-2/myHeap"
)

func ReadDishes() (dishes myHeap.MyHeap) {
	var dishCount int
	fmt.Scan(&dishCount)

	for i := 0; i < dishCount; i++ {
		var temp int
		fmt.Scan(&temp)
		heap.Push(&dishes, temp)
	}
	return dishes
}
