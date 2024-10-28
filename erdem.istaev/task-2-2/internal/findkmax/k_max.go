package findkmax

import (
	"container/heap"

	"erdem.istaev/task-2-2/internal/structure"
)

func FindKMax(k int, dishes []int, n int) int {
	h := &structure.IntHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, dishes[i])
	}

	for i := k; i < n; i++ {
		if dishes[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, dishes[i])
		}
	}

	return (*h)[0]
}
