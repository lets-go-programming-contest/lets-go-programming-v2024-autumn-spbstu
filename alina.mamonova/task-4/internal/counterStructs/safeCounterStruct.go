package counterStructs

import "sync"

type SafeCounter struct {
	Count int
	Mutex sync.Mutex
}

func (counter *SafeCounter) Increment() {
	counter.Mutex.Lock()
	counter.Count++
	counter.Mutex.Unlock()
}

func (counter *SafeCounter) Value() int {
	counter.Mutex.Lock()
	defer counter.Mutex.Unlock()
	return counter.Count
}
