package test

import (
	"fmt"
	"sync"

	"github.com/Madyarov-Gleb/task-4/internal/counter"
)

func Test() {
	counterExample := counter.Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go counterExample.Enter(&wg)
		fmt.Println("man", i, "enter")
	}
	wg.Wait()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go counterExample.Exit(&wg)
		fmt.Println("man", i, "exit")
	}
	wg.Wait()
	fmt.Println(counterExample.GetValue())
}

func TestUnsafe() {
	counterExample := counter.Counter{}
	channel := make(chan int)
	var id int
	for i := 0; i < 10; i++ {
		go counterExample.EnterUnsafe(i, channel)
	}
	for i := 0; i < 10; i++ {
		id = <-channel
		fmt.Println("man", id, "enter")
	}
	for i := 0; i < 3; i++ {
		go counterExample.ExitUnsafe(i, channel)
	}
	for i := 0; i < 3; i++ {
		id = <-channel
		fmt.Println("man", id, "exit")
	}
	fmt.Println(counterExample.GetValue())
}
