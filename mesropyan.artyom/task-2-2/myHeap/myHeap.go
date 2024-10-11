package MyHeap

type MyHeap struct {
	Elems []int
}

func (h MyHeap) Push(x int) {
	h.Elems = append(h.Elems, x)
}

func (h MyHeap) Pop() int {
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
	if h.Elems[i] < h.Elems[j] {
		return true
	}
	return false
}

func (h MyHeap) Swap(i, j int) {
	temp := h.Elems[i]
	h.Elems[i] = h.Elems[j]
	h.Elems[j] = temp
}
