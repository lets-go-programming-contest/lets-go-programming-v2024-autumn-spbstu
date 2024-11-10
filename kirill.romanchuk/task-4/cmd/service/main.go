package main

import (
	"github.com/kirill.romanchuk/task-4/structures/performance"
)

func main() {
	perm := performance.CreatePerformance(3, 4)
	perm.DisplaySeats()
}
