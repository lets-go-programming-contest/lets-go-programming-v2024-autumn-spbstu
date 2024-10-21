package maxk

import (
	"container/heap"
	"errors"
	"github.com/sssidkn/task-2-2/pkg/iheap"
)

func Find(dishes []int, k int) (int, error) {
	if k < 0 || k > len(dishes) {
		return -1, errors.New("invalid k")
	}

	h := &iheap.IntHeap{}
	heap.Init(h)

	for _, v := range dishes {
		heap.Push(h, v)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0], nil
}
