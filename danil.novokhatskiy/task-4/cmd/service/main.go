package main

import (
	"fmt"
	"sync"

	"github.com/katagiriwhy/task-4/internal"
)

func main() {
	counter := internal.SaveCounter{}
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
