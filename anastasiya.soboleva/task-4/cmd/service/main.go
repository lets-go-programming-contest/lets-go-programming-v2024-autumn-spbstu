package main

import (
	"anastasiya.soboleva/task-4/internal/exhibit"
	"anastasiya.soboleva/task-4/internal/simulator"
)

func main() {
	visitorCounts := []int{10, 15}

	exhibits := []*exhibit.Exhibit{
		exhibit.NewExhibit("African Safari"),
		exhibit.NewExhibit("Underwater World"),
	}

	simulator.RunSafeSimulation(exhibits, visitorCounts)
	simulator.RunUnsafeSimulation(exhibits, visitorCounts)
}
