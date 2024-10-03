package main

import (
	"container/heap"
	"errors"
	"fmt"
	"log"
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
	fmt.Println("Enter n:")
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(errors.New("incorrect data"))
	}

	fmt.Println("Enter ai:")

	numbers := make([]int, n)
	i := 0
	var ai int
	for ; i < n; i++ {
		_, err = fmt.Scan(&ai)
		if err != nil {
			log.Fatal(errors.New("incorrect data"))
		}
		numbers[i] = ai
	}

	fmt.Println("Enter k:")
	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		log.Fatal(errors.New("incorrect data"))
	}

	kMax, err := findKMax(&numbers, k)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kMax)
}

func findKMax(nums *[]int, k int) (int, error) {
	if k <= 1 || k > len((*nums)) {
		return 0, errors.New("incorrect k")
	}

	h := &IntHeap{}
	heap.Init(h)
	for _, num := range *nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0], nil
}
