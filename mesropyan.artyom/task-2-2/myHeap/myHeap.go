package myHeap

type MyHeap []int

func (h *MyHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MyHeap) Pop() interface{} {
	len := h.Len()
	res := (*h)[len-1]
	*h = (*h)[:len-1]
	return res
}

func (h *MyHeap) Len() int {
	return len(*h)
}

func (h *MyHeap) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

func (h *MyHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
