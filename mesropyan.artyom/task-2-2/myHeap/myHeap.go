package myHeap

type MyHeap struct {
	Elems []int
}

func (h MyHeap) Push(x any) {
	if elem, ok := x.(int); ok {
		h.Elems = append(h.Elems, elem)
	}
}

func (h MyHeap) Pop() any {
	if h.Len() != 0 {
		result := h.Elems[h.Len()-1]
		h.Elems = h.Elems[:h.Len()-1]
		return result
	}
	return 0
}

func (h MyHeap) Len() int {
	return len(h.Elems)
}

func (h MyHeap) Less(i, j int) bool {
	return h.Elems[i] < h.Elems[j]
}

func (h MyHeap) Swap(i, j int) {
	temp := h.Elems[i]
	h.Elems[i] = h.Elems[j]
	h.Elems[j] = temp
}
