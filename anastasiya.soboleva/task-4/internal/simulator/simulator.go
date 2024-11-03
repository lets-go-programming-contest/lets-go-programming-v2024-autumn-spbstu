package simulator

import (
	"errors"
	"fmt"
	"sync"

	"anastasiya.soboleva/task-4/internal/exhibit"
)

func RunSafeSimulation(exhibits []*exhibit.Exhibit, visitorCounts []int) error {
	if len(exhibits) != len(visitorCounts) {
		return errors.New("количество выставок не совпадает с количеством посетителей")
	}

	var wg sync.WaitGroup
	for i, e := range exhibits {
		wg.Add(1)
		go func(ex *exhibit.Exhibit, count int) {
			defer wg.Done()
			ex.SimulateVisitorSafe(count)
		}(e, visitorCounts[i])
	}
	wg.Wait()
	fmt.Println("Результаты симуляции с синхронизацией:")
	for _, e := range exhibits {
		e.ShowInfo()
	}
	return nil
}

func RunUnsafeSimulation(exhibits []*exhibit.Exhibit, visitorCounts []int) error {
	if len(exhibits) != len(visitorCounts) {
		return errors.New("количество выставок не совпадает с количеством посетителей")
	}

	var wg sync.WaitGroup
	for i, e := range exhibits {
		wg.Add(1)
		go func(ex *exhibit.Exhibit, count int) {
			defer wg.Done()
			ex.SimulateVisitorUnsafe(count)
		}(e, visitorCounts[i])
	}
	wg.Wait()
	fmt.Println("Результаты симуляции без синхронизации:")
	for _, e := range exhibits {
		e.ShowInfo()
	}
	return nil
}
