package dishHeap

type dishHeap []int

func (h dishHeap) Len() int           { return len(h) }
func (h dishHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h dishHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *dishHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *dishHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
