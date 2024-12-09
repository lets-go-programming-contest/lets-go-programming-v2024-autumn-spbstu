package heapOfDishes

type HeapOfDishes []int

func (h HeapOfDishes) Len() int {
	return len(h)
}

func (h HeapOfDishes) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h HeapOfDishes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapOfDishes) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *HeapOfDishes) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
