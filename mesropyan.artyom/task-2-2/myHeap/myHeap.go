package myHeap

type MyHeap []int

func (h *MyHeap) Push(x interface{}) {
	if elem, ok := x.(int); ok {
		*h = append(*h, elem)
	}
}

func (h *MyHeap) Pop() interface{} {
	if h.Len() != 0 {
		result := (*h)[(h).Len()-1]
		*h = (*h)[:h.Len()-1]
		return result
	}
	return 0
}

func (h MyHeap) Len() int {
	return len(h)
}

func (h MyHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MyHeap) Swap(i, j int) {
	temp := h[i]
	h[i] = h[j]
	h[j] = temp
}
