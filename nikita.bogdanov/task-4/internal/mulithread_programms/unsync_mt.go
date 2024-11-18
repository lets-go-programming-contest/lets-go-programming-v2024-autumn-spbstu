package mulithread_programms

import (
	"fmt"
	"sync"
	"time"
)

type TurnstileCounterUnsync struct {
	count int
}

func (tc *TurnstileCounterUnsync) Increment() {
	tc.count++
}

func (tc *TurnstileCounterUnsync) GetCount() int {
	return tc.count
}

func UnsyncMT() {
	var wg sync.WaitGroup
	counter := TurnstileCounterUnsync{}

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
