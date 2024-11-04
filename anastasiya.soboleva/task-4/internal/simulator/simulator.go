package simulator

import (
	"errors"
	"fmt"
	"sync"

	"anastasiya.soboleva/task-4/internal/exhibit"
)

const visitorBatchSize = 1000

func runSimulation(exhibits []*exhibit.Exhibit, visitorCounts []int, safe bool) error {
	if len(exhibits) != len(visitorCounts) {
		return errors.New("количество выставок не совпадает с количеством посетителей")
	}
	var wg sync.WaitGroup
	for i, e := range exhibits {
		count := visitorCounts[i]
		for j := 0; j < count; j += visitorBatchSize {
			batchSize := visitorBatchSize
			if j+batchSize > count {
				batchSize = count - j
			}
			wg.Add(1)
			go func(ex *exhibit.Exhibit, batch int) {
				defer wg.Done()
				if safe {
					ex.SimulateVisitorSafe(batch)
				} else {
					ex.SimulateVisitorUnsafe(batch)
				}
			}(e, batchSize)
		}
	}
	wg.Wait()
	if safe {
		fmt.Println("Результаты симуляции с синхронизацией:")
	} else {
		fmt.Println("Результаты симуляции без синхронизации:")
	}
	for _, e := range exhibits {
		e.ShowInfo()
	}
	return nil
}

func RunSafeSimulation(exhibits []*exhibit.Exhibit, visitorCounts []int) error {
	return runSimulation(exhibits, visitorCounts, true)
}

func RunUnsafeSimulation(exhibits []*exhibit.Exhibit, visitorCounts []int) error {
	return runSimulation(exhibits, visitorCounts, false)
}
