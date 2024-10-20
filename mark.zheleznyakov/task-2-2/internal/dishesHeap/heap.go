package dishesHeap

import (
	"container/heap"
)

type DishesHeap []int

func (h DishesHeap) Len() int {
	return len(h)
}

func (h DishesHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h DishesHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DishesHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *DishesHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindDish(h *DishesHeap, k int) int {
	var result int
	for _ = range k {
		result = heap.Pop(h).(int)
	}
	return result
}
