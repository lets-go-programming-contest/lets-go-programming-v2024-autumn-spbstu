package main

import (
	"fmt"
	"sync"

	"github.com/mrqiz/task-4/internal/parking"
	"github.com/mrqiz/task-4/internal/simulation"
)

func main() {
	park := parking.NewParking(5)
	var wg sync.WaitGroup
	fmt.Printf("okay so there are %d slots in the parking\n", park.Capacity())
	simulation.Simulate(park, &wg)
}
