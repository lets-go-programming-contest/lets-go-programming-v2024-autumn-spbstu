package heap

type Heap []int

func (h *Heap) Len() int {
	return len(*h)
}
func (h *Heap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}
func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
