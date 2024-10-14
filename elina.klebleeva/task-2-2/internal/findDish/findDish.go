package findDish

import (
	"container/heap"
	"task-2-2/pkg/heapInt"
)

func FindDish(dishNum int, dishes heapInt.HeapInt) any {

	tempDishes := make(heapInt.HeapInt, len(dishes))
	copy(tempDishes, dishes)
	heap.Init(&tempDishes)

	for i := 0; i < dishNum-1; i++ {
		heap.Pop(&tempDishes)
	}

	return heap.Pop(&tempDishes)
}
