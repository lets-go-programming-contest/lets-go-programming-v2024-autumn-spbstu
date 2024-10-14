package meal

import (
	"container/heap"

	hp "anastasiya.soboleva/task-2-2/internal/heap"
)

func FindKMeal(meals []int, k int) int {
	h := &hp.Heap{}
	heap.Init(h)

	for _, meal := range meals {
		heap.Push(h, meal)
	}

	var kMeal int
	for i := 0; i < k; i++ {
		kMeal = heap.Pop(h).(int)
	}

	return kMeal
}
