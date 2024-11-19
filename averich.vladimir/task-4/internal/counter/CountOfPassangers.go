package counter

import (
	"sync"
)

type TurnstileCounter struct {
	mu    sync.Mutex
	count int
}

func NewTurnstileCounter(count int) *TurnstileCounter {
	return &TurnstileCounter{
		count: count,
	}
}

func (tc *TurnstileCounter) PassThrough() {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	tc.count++
}

func (tc *TurnstileCounter) GetCount() int {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	return tc.count
}