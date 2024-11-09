package main

import (
	"sync"

	"github.com/mrqiz/task-4/internal/parking"
	"github.com/mrqiz/task-4/internal/simulation"
)

func main() {
	park := parking.NewParking()
	var wg sync.WaitGroup
	simulation.Simulate(park, &wg)
	simulation.UnsafeSimulate(park, &wg)
}
