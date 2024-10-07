package main

import (
	"container/heap"
	"fmt"
	"os"
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
	_, err := fmt.Scan(&N)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dishes := make([]int, N)
	for i := 0; i < N; i++ {
		_, err = fmt.Scan(&temp)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		dishes[i] = temp
	}

	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		fmt.Println(err)
	}
	h := &IntHeap{}
	heap.Init(h)
	if k < N && k > 0 {
		for _, v := range dishes {
			heap.Push(h, v)
			if h.Len() > k {
				heap.Pop(h)
			}
		}
		fmt.Println((*h)[0])
	} else {
		fmt.Println(-1)
	}
}
