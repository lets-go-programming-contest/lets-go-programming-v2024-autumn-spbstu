package meatHeap

import (
	"container/heap"
)

type MeatHeap []int

func (h MeatHeap) Len() int {
	return len(h)
}

func (h MeatHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MeatHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MeatHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MeatHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindDish(h *MeatHeap, k int) int {
	var result int
	for range k {
		result = heap.Pop(h).(int)
	}
	return result
}
