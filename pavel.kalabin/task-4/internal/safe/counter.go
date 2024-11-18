package safe 

import "sync"

type Counter struct {
	Val int
	RWMutex sync.RWMutex
}

func (c *Counter) Count() {
	c.RWMutex.Lock()
	defer c.RWMutex.Unlock()
	c.Val = c.Val + 1 // can be atomic.AddInt32
}

func (c *Counter) Value() int {
	c.RWMutex.RLock()
	defer c.RWMutex.RUnlock()
	return c.Val
}

