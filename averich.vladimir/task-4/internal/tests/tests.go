package tests

import (
	"fmt"
	"sync"

	"task-4/internal/counter"
	"task-4/internal/unsafeCounter"
)

const GoroutinsAmount = 1000

func ExecuteSafeTest() {
	fmt.Println("Started safe test\n")
	var waitGroup sync.WaitGroup

	counter := counter.NewTurnstileCounter(0)

	for i := 0; i < GoroutinsAmount; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			counter.PassThrough()
		}()
	}
	waitGroup.Wait()

	fmt.Println("Ended safe test\n")
	fmt.Println("Amount of passenger, passed control:", counter.GetCount())
}

func ExecuteUnSafeTest() {
	fmt.Println("Started unsafe test\n")
	var waitGroup sync.WaitGroup

	counter := unsafeCounter.NewTurnstileCounter(0)

	for i := 0; i < GoroutinsAmount; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			counter.PassThrough()
		}()
	}
	waitGroup.Wait()

	fmt.Println("Ended unsafe test")
	fmt.Println("Amount of passenger, passed control:", counter.GetCount())
}