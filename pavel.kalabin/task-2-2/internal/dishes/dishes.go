package dishes

import "container/heap"
/*
type Interface interface {
    sort.Interface
    Push(x interface{}) // вставляет x как элемент Len()
    Pop() interface{} // удаляет и возвращает элемент Len() - 1
}

type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
*/
type DishesHeap []int

func (dh *DishesHeap) Push(value interface{}) {
    a, ok := value.(int)
    if ok {
        *dh = append(*dh, a)
    }
}

func (dh *DishesHeap) Pop() (value interface{}) {
    oldH := *dh
    n := len(oldH)
    value = oldH[n - 1]
    *dh = oldH[0:n-1]
    return 
}

func (dh *DishesHeap) Len() int {
    return len(*dh)
}

func (dh *DishesHeap) Less(i, j int) bool {
    return (*dh)[i] > (*dh)[j]
}

func (dh *DishesHeap) Swap(i, j int) {
    (*dh)[i], (*dh)[j] = (*dh)[j], (*dh)[i]
}

func FindKMax(h *DishesHeap, k int) (kMax int) {
    for range k {
        kMax = heap.Pop(h).(int)
    }
    return
}

