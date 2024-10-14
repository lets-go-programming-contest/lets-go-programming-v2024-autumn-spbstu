package heapInt

type HeapInt []int

func (h *HeapInt) Len() int {
	return len((*h))
}

func (h *HeapInt) Less(i int, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *HeapInt) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HeapInt) Push(x interface{}) {
	a, ok := x.(int)
	if ok {
		*h = append(*h, a)
	} else {
		panic("incorrect type for HeatPush")
	}

}

func (h *HeapInt) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
