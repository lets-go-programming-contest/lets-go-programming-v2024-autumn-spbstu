package main

import (
	"fmt"
	"sync"

	"github.com/Piyavva/task-4/internal/turnstile"
)

func main() {
	tur := turnstile.Turnstile{}
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tur.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Mutex: %d\n", tur.GetValue())
	utur := turnstile.UnsafeTurnstile{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			utur.Increment()
		}()
	}
	wg.Wait()
	fmt.Printf("Without Mutex: %d\n", utur.GetValue())

}
