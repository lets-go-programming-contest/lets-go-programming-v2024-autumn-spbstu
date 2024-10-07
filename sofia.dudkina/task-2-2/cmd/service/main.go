package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func main() {
	var N, temp int
	fmt.Scan(&N)

	dishes := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&temp)
		dishes[i] = temp
	}

	var k int
	fmt.Scan(&k)

	h := &IntHeap{}
	heap.Init(h)

	for _, v := range dishes {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	fmt.Println((*h)[0])
}
