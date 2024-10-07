package main

import (
	"container/heap"
	"fmt"
	"log"
)

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(i, j int) bool { return h[i] < h[j] }

func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKDish(arr *[]int, k int) {
	if k > len(*arr) || k <= 1 {
		log.Fatal("Invalid input")
	}
	h := &Heap{}
	heap.Init(h)
	for _, v := range *arr {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	fmt.Println("Result:", (*h)[0])
}

func main() {
	fmt.Print("Enter the number of dishes: ")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		log.Fatal(err)
	}
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		var tmp int
		_, err := fmt.Scan(&tmp)
		if err != nil || tmp > 10000 || tmp < -10000 {
			fmt.Println("Invalid input")
		}
		arr[i] = tmp
	}
	fmt.Println("Enter k value: ")
	var k int
	_, err = fmt.Scan(&k)
	if err != nil {
		log.Fatal(err)
	}
	findKDish(&arr, k)
}
