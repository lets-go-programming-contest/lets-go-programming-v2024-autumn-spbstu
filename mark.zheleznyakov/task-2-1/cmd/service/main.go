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
	depts := input.ReadInt()

	for _ = range depts {
		minTemp := initialMinTemp
		maxTemp := initialMaxTemp
		workersPerDept := input.ReadInt()
    
    for _ = range workersPerDept {
      condition, tVal := input.ReadCondition()
      
			if condition == ">=" && tVal > minTemp {
				minTemp = tVal
			} else if condition == "<=" && tVal < maxTemp {
				maxTemp = tVal
			}
		}

    if minTemp > maxTemp {
			fmt.Println(-1)
		} else {
			fmt.Println(minTemp)
		}
	}
}

