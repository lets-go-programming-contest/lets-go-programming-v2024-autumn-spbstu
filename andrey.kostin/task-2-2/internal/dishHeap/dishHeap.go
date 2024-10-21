package dishHeap

type DishHeap []int

func (h *DishHeap) Len() int           { return len(*h) }
func (h *DishHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *DishHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *DishHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *DishHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
