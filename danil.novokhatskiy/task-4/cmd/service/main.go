package main

import (
	"fmt"
	"sync"
)

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

func main() {
	counter := SaveCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Decrement()
		}()
	}
	wg.Wait()
	fmt.Printf("Final Counter = %d\n", counter.Value())
}
