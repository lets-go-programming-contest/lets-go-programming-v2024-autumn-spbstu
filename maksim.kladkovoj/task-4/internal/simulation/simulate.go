package simulation

import (
	"fmt"
	"sync"

	"github.com/Mmmakskl/task-4/internal/structures"
)

func Simulate(in, out int) {
	fmt.Println("Safe simulation")

	counter := &structures.TrainCounter{}
	var wg sync.WaitGroup

	for i := 0; i < in; i++ {
		wg.Add(1)
		go counter.Arrive(&wg)
	}
	wg.Wait()

	fmt.Printf("%d train have arrived\n", counter.GetArrival())

	for j := 0; j < out; j++ {
		wg.Add(1)
		go counter.Depart(&wg)
	}
	wg.Wait()

	fmt.Printf("%d train have departed \n\n", counter.GetDeparted())
}
