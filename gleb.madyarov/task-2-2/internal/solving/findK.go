package solving

import (
	"container/heap"

	"github.com/Madyarov-Gleb/task-2-2/internal/iheap"
)

func FindKMax(nums *[]int, k int) int {
	h := &iheap.IntHeap{}
	heap.Init(h)
	for _, num := range *nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}
