package lunch

import (
	"container/heap"
	"fmt"
	myHeap "task-2-2/internal/heap"
	"task-2-2/internal/reader"
)

type ConsoleLunch struct{}

func NewConsoleLunch() *ConsoleLunch {
	return &ConsoleLunch{}
}

func findKthElement(slice []int, k int) int {
	data := &myHeap.IntHeap{}
	heap.Init(data)

	for _, e := range slice {
		heap.Push(data, e)
	}

	for i := 0; i < k-1; i++ {
		heap.Pop(data)
	}

	return heap.Pop(data).(int)
}

func (c *ConsoleLunch) Run() {
	fmt.Println("Поиск k-го наибольшего числа массива")
	fmt.Println("'выход' - для завершения")

	for {
		sliceSize, exit, err := reader.ReadSliceSize()
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		slice, exit, err := reader.ReadSliceElements(sliceSize)
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		k, exit, err := reader.ReadKthElement(sliceSize)
		if exit {
			return
		}
		if err != nil {
			fmt.Printf("%v\n", err)
			continue
		}

		fmt.Printf("Результат: %d\n", findKthElement(slice, k))
	}
}
