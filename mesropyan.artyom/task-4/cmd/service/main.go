package main

import (
	"sync"

	"github.com/artem6554/task-4/queue"
)

func main() {
	var queue queue.Queue
	var wGroup sync.WaitGroup
	for i := 0; i < 5; i++ {
		wGroup.Add(1)
		go func() {
			defer wGroup.Done()
			queue.Push(i)
			queue.PrintAll()
		}()
	}

	for i := 0; i < 5; i++ {
		wGroup.Add(1)
		go func() {
			defer wGroup.Done()
			queue.Pop()
		}()
	}
	wGroup.Wait()
}
