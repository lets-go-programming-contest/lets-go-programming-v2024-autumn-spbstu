package main

import (
	"log"

	"anastasiya.soboleva/task-4/internal/exhibit"
	"anastasiya.soboleva/task-4/internal/simulator"
)

func main() {
	visitorCounts := []int{10, 15}

	exhibits := []*exhibit.Exhibit{
		exhibit.NewExhibit("Древние цивилизации"),
		exhibit.NewExhibit("Секреты океана"),
	}

	if err := simulator.RunSafeSimulation(exhibits, visitorCounts); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	if err := simulator.RunUnsafeSimulation(exhibits, visitorCounts); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}
