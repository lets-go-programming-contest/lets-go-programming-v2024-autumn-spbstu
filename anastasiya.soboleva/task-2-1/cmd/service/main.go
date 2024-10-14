package main

import (
	"fmt"

	"anastasiya.soboleva/task-2-1/internal/reader"
	"anastasiya.soboleva/task-2-1/internal/temperature"
)

func main() {
	n := reader.ReadDepartmentCount()

	for i := 0; i < n; i++ {
		k := reader.ReadEmployeeCount()

		tracker := temperature.NewTracker()

		for j := 0; j < k; j++ {
			operation, value := reader.ReadTemperatureOperation()
			result := tracker.ProcessTemperature(operation, value)
			fmt.Println(result)
		}
	}
}
