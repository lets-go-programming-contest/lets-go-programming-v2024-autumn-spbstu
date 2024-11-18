package metro

import (
	"sync"
)

type Turnstile struct {
	mutex sync.Mutex
	count int
}

func (t *Turnstile) Enter(wg *sync.WaitGroup) {
	defer wg.Done()
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.count++
}

func (t *Turnstile) RawEnter() {
	t.count++
}

func (t *Turnstile) GetCount() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.count
}

type TurnstileCollection struct {
	turnstiles []*Turnstile
}

func NewTurnstileCollection(num int) *TurnstileCollection {
	collection := &TurnstileCollection{}
	for i := 0; i < num; i++ {
		collection.turnstiles = append(collection.turnstiles, &Turnstile{})
	}
	return collection
}

func (tc *TurnstileCollection) GetCount() int {
	totalCount := 0
	for _, turnstile := range tc.turnstiles {
		totalCount += turnstile.GetCount()
	}
	return totalCount
}

func (tc *TurnstileCollection) ResetCounts() {
	for _, turnstile := range tc.turnstiles {
		turnstile.mutex.Lock()
		turnstile.count = 0
		turnstile.mutex.Unlock()
	}
}
