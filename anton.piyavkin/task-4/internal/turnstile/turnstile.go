package turnstile

import "sync"

type Turnstile struct {
	people int
	mut    sync.Mutex
}

func (t *Turnstile) Increment() {
	t.mut.Lock()
	t.people++
	t.mut.Unlock()
}

func (t *Turnstile) GetValue() int {
	return t.people
}
