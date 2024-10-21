package main

import (
	"fmt"
	"task-2-1/internal/analyzer"
	"task-2-1/internal/temperature"
)

func main() {
	n := analyzer.AnalyzeDeptCount()

	for i := 0; i < n; i++ {

		k := analyzer.AnalyzeWorkersCount()

		tracker := temperature.NewTracker()

		for j := 0; j < k; j++ {
			operation, value := analyzer.AnalyzeTemp()
			result := tracker.ProgTemperature(operation, value)
			fmt.Println(result)
		}
	}
}
