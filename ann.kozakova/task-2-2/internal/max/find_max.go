package max

import (
	"container/heap"
	"errors"

	"github.com/nutochk/task-2-2/pkg/int_heap"
)

func FindKMax(nums *[]int, k int) (int, error) {
	if k <= 1 || k > len((*nums)) {
		return 0, errors.New("incorrect k")
	}

	h := &int_heap.IntHeap{}
	heap.Init(h)
	for _, num := range *nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0], nil
}
