package main

import (
	"fmt"
)

type Heap []int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	size := len(*h)
	x := (*h)[size - 1]
    (*h)[size - 1] = 0
	*h = (*h)[:size - 1]
	return x
}

func main() {
    var h Heap
    h.Push(1)
    h.Push(2)
    fmt.Println(h)
    h.Pop()
    fmt.Println(h)
    h.Pop()
}