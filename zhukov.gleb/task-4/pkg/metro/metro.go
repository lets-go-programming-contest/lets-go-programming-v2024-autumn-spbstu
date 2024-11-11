package metro

import "sync"

type Visitors struct {
	Quantity map[int]struct{}
	mu       *sync.RWMutex
}

type VisitorRepo interface {
	Register(id int)
	GetCount() int
	Simulator(cntVisitors int, out chan<- string)
}
