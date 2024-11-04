package main

import (
	"anastasiya.soboleva/task-4/internal/exhibit"
	"anastasiya.soboleva/task-4/internal/simulator"
	"log"
)

func main() {
	visitorCounts := []int{
		20000, 20000,
	}
	exhibits1 := []*exhibit.Exhibit{
		exhibit.NewExhibit("Древние цивилизации"),
		exhibit.NewExhibit("Тайны океана"),
	}
	if err := simulator.RunSafeSimulation(exhibits1, visitorCounts); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
	exhibits2 := []*exhibit.Exhibit{
		exhibit.NewExhibit("Древние цивилизации"),
		exhibit.NewExhibit("Тайны океана"),
	}
	if err := simulator.RunUnsafeSimulation(exhibits2, visitorCounts); err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}
