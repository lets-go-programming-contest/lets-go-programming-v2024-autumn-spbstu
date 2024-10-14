package getDish

import (
	"container/heap"

	"github.com/Piyavva/task-2-2/internal/myHeap"
)

func GetKDish(dishes []int, k int) int {
    h := &myHeap.Heap{}
    heap.Init(h)
    for i := 0; i < k; i++ {
        heap.Push(h, dishes[i])
    }
    for i := k; i < len(dishes); i++ {
        if dishes[i] > (*h)[0] {
            heap.Pop(h)
            heap.Push(h, dishes[i])
        }
    }
    return (*h)[0]
}