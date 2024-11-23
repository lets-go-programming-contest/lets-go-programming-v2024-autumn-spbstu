package internal

import "sync"

type SaveCounter struct {
	Mutex   sync.Mutex
	Counter int
}

func (c *SaveCounter) Increment() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Counter++
}

func (c *SaveCounter) Decrement() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Counter--
}

func (c *SaveCounter) Value() int {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	return c.Counter
}
