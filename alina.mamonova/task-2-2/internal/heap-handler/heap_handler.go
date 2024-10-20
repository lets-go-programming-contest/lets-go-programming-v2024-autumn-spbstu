package heap_handler

import (
	"container/heap"
	"github.com/hahapathetic/task-2-2/internal/input"
	"github.com/hahapathetic/task-2-2/internal/table-heap"
)

const (
	minAIBound = -10000
	maxAIBound = 10000
)

func FillHeap(size int, table *table_heap.TableHeap) {
	heap.Init(table)
	for i := 0; i < size; i++ {
		ai := input.ProcessNumInput("Enter the ai: ", minAIBound, maxAIBound)
		heap.Push(table, ai)
	}
}

func ProcessHeap(size int, table *table_heap.TableHeap) int {
	k := input.ProcessNumInput("Enter the k: ", 1, size)
	var result int = 0
	for i := 0; i < k; i++ {
		result = heap.Pop(table).(int)
	}
	return result
}
