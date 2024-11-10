package main

import (
	"container/heap"
	"fmt"
	"log"

	"github.com/katagiriwhy/task-2-2/internal"
)

func findKDish(arr *[]int, k int) {
	if k > len(*arr) || k <= 1 {
		log.Fatal("Invalid input")
	}
	h := &internal.Heap{}
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
	arr, k, err := internal.ReadData()
	if err != nil {
		log.Fatal(err)
	} else {
		findKDish(&arr, k)
	}
}
