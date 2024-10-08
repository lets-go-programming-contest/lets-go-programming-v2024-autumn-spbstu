package main

import (
	"fmt"
	"github.com/mrqiz/task-2-1/internal/input"
)

func main() {
	depts := input.ReadInt("departments count")
	workersPerDept := input.ReadInt("workers per department")
	
	fmt.Printf("got depts: %d\n", depts)
	fmt.Printf("got workersPerDept: %d\n", workersPerDept)
}

