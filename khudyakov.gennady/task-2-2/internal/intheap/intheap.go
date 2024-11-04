package intheap

type IntHeap []int

func (heap IntHeap) Len() int {
	return len(heap)
}

func (heap IntHeap) Less(i, j int) bool {
	return heap[i] > heap[j]
}

func (heap *IntHeap) Swap(i, j int) {
	buffer := (*heap)[i]
	(*heap)[i] = (*heap)[j]
	(*heap)[j] = buffer
}

func (heap *IntHeap) Push(x any) {
	*heap = append(*heap, x.(int))
}

func (heap *IntHeap) Pop() any {
	result := (*heap)[heap.Len()-1]
	*heap = (*heap)[:heap.Len()-1]
	return result
}
