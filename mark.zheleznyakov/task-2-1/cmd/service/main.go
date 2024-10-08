package main

import (
	"fmt"
	"github.com/mrqiz/task-2-1/internal/input"
)

var (
	initialMinTemp = 15
	initialMaxTemp = 30
)

func main() {
	depts := input.ReadInt("departments count")

	for _ = range depts {
		workersPerDept := input.ReadInt("workers per department")
    
    for _ = range workersPerDept {
      condition, tVal := input.ReadCondition()
			fmt.Println(condition, tVal)
		}
	}
}

