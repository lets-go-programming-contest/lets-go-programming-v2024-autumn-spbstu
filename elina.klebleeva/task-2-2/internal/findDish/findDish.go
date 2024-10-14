package findDish

import (
	"container/heap"
	"task-2-2/pkg/heapInt"
)

func FindDish(dishNum int, dishes heapInt.HeapInt) any {
	for i := 0; i < dishNum-1; i++ {
		heap.Pop(&dishes)
	}

	return heap.Pop(&dishes)
}
