package tests

import (
	"fmt"
	"github.com/hahapathetic/task-4/internal/counterStructs"
	"sync"
)

const goroutinesNum = 50

func TestSafeStruct() {
	var wg sync.WaitGroup
	counter := &counterStructs.SafeCounter{}

	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			counter.Increment()
			fmt.Printf("[Sync] Goroutine %d, value: %d\n", id, counter.Value())
		}(i)
	}

	wg.Wait()

	fmt.Printf("Final number: %d\n", counter.Value())
}

func TestUnsafeStruct() {

	var wg sync.WaitGroup
	counter := &counterStructs.UnsafeCounter{}

	for i := 0; i < goroutinesNum; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			counter.Increment()
			fmt.Printf("[Without sync] Goroutine %d, value: %d\n", id, counter.Value())
		}(i)
	}

	wg.Wait()

	fmt.Printf("Final number: %d\n", counter.Value())
}
