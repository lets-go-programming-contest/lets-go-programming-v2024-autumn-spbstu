package main

import (
	"container/heap"
	"fmt"
	"os"
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

func getKDish(dishes []int, k int) int {
    h := &Heap{}
	heap.Init(h)
    for i := 0; i < k; i++ {
        heap.Push(h, dishes[i])
    }
    for i := k; i < len(dishes); i++ {
        if dishes[i] > (*h)[0] {
            heap.Pop(h)
            heap.Push(h, dishes[i])
        }
    }
    return (*h)[0]
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
        os.Exit(1)
	}
	var dish int
    dishes := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Scan(&dish)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
            os.Exit(1)
		}
		dishes[i] = dish
	}
    var k int
	_, err = fmt.Scan(&k)
    if err != nil || k < 1 || k > n {
        fmt.Fprint(os.Stderr, "invalid k\n")
        os.Exit(1)
    }
    res := getKDish(dishes, k)
	fmt.Println(res)
}