package main

import (
	"fmt"

	"github.com/Koshsky/task-4/internal/metro"
)

func main() {
	numTurnstiles := 5
	numPeople := 100
	turnstileCollection := metro.NewTurnstileCollection(numTurnstiles)

	fmt.Println("Запускаем горутины для симуляции с синхронизацией")
	metro.SimulateWithSync(turnstileCollection, numPeople)

	turnstileCollection.ResetCounts()

	fmt.Println("Запускаем горутины для симуляции без синхронизации")
	metro.SimulateWithoutSync(turnstileCollection, numPeople)
}
