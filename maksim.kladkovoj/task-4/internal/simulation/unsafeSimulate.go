package simulation

import (
	"fmt"

	"github.com/Mmmakskl/task-4/internal/structures"
)

func UnsafeSimulate(in, out int) {
	fmt.Println("Unsafe simulation")

	counter := &structures.TrainCounter{}

	for i := 0; i < in; i++ {
		go counter.ArriveUnsafe()
	}
	fmt.Printf("%d train have arrived \n", counter.GetArrival())
	for j := 0; j < out; j++ {
		go counter.DepartUnsafe()
	}
	fmt.Printf("%d train have departed \n\n", counter.GetDeparted())
}
