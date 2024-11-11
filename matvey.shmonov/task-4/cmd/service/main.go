package main

import (
	"fmt"

	"github.com/Koshsky/task-4/internal/metro"
)

func main() {
	numTurnstiles := 5
	numPeople := 100
	turnstileCollection := metro.NewTurnstileCollection(numTurnstiles)

	fmt.Println("Running goroutines for synchronized simulation")
	metro.SimulateWithSync(turnstileCollection, numPeople)

	turnstileCollection.ResetCounts()

	fmt.Println("Running goroutines for simulation without synchronization")
	metro.SimulateWithoutSync(turnstileCollection, numPeople)
}
