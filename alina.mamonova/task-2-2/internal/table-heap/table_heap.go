package table_heap

type TableHeap []int

func (h TableHeap) Len() int           { return len(h) }
func (h TableHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h TableHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *TableHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *TableHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
