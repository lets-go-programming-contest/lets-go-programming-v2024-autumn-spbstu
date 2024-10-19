package findDish

import (
	"container/heap"

	"github.com/IDevFrye/task-2-2/internal/dishHeap"
)

func FindDish(prefs *[]int, k int) int {
	dishes := &dishHeap.DishHeap{}
	heap.Init(dishes)
	for _, pref := range *prefs {
		heap.Push(dishes, pref)
	}
	var result int
	for i := 0; i < k; i++ {
		result = heap.Pop(dishes).(int)
	}
	return result
}
