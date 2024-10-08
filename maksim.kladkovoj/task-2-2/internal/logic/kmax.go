package logic

import (
	"container/heap"

	"github.com/Mmmakskl/task-2-2/internal/structure"
)

func KMax(k int, rating *structure.IntHeap) int {
	for i := 1; i < k; i++ {
		heap.Pop(rating)
	}
	return (*rating)[0]
}
