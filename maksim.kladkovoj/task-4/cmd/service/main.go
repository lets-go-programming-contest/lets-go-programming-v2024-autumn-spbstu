package main

import (
	"github.com/Mmmakskl/task-4/internal/simulation"
)

func main() {
	in, out := 100, 85
	simulation.Simulate(in, out)
	simulation.UnsafeSimulate(in, out)
}
