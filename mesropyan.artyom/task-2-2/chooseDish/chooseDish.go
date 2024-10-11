package chooseDish

import (
	"container/heap"
	"fmt"

	"github.com/artem6554/task-2-2/myHeap"
)

func ChooseDish(dishes myHeap.MyHeap) int {
	var num int
	fmt.Scan(&num)
	for i := 0; i < num-1; i++ {
		heap.Pop(dishes)
	}
	result := heap.Pop(dishes).(int)
	return result
}
