package heap_handler

import (
	"bufio"
	"container/heap"
	"github.com/hahapathetic/task-2-2/internal/input"
	"github.com/hahapathetic/task-2-2/internal/table-heap"
)

const (
	minAIBound = -10000
	maxAIBound = 10000
)

func FillHeap(size int, table *table_heap.TableHeap, reader *bufio.Reader) {
	heap.Init(table)
	values := input.ProcessHeapInput(minAIBound, maxAIBound, size, reader)
	for i := 0; i < size; i++ {
		ai := values[i]
		heap.Push(table, ai)
	}
}

func ProcessHeap(size int, table *table_heap.TableHeap, k int) int {
	var result int = 0
	for i := 0; i < k; i++ {
		result = heap.Pop(table).(int)
	}
	return result
}
