package visitorCounter

import (
	"sync"
)

type VisitorCounter struct {
	count int
	mu    sync.Mutex
}

func NewVisitorCounter() *VisitorCounter {
	return &VisitorCounter{count: 0}
}

func (vc *VisitorCounter) Increment() {
	vc.mu.Lock()
	vc.count++
	vc.mu.Unlock()
}

func (vc *VisitorCounter) UnsafeIncrement() {
	vc.count++
}

func (vc *VisitorCounter) GetCount() int {
	vc.mu.Lock()
	defer vc.mu.Unlock()
	return vc.count
}
