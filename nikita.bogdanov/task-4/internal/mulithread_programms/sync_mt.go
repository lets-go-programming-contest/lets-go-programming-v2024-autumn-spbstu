package mulithread_programms

import (
	"fmt"
	"sync"
	"time"
)

type TurnstileCounterSync struct {
	mu    sync.Mutex
	count int
}

func (tc *TurnstileCounterSync) Increment() {
	tc.mu.Lock()
	tc.count++
	tc.mu.Unlock()
}

func (tc *TurnstileCounterSync) GetCount() int {
	tc.mu.Lock()
	defer tc.mu.Unlock()
	return tc.count
}

func SyncMT() {
	var wg sync.WaitGroup
	counter := TurnstileCounterSync{}

	numPeople := 1000
	wg.Add(numPeople)

	for i := 0; i < numPeople; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
			time.Sleep(time.Millisecond)
		}()
	}

	wg.Wait()
	fmt.Printf("Total people passed through turnstile: %d\n", counter.GetCount())
}
