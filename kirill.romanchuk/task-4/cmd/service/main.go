package main

import (
	"sync"
	"time"

	"github.com/kirill.romanchuk/task-4/structures/performance"
)

func main() {
	var mutex sync.Mutex
	perm := performance.CreatePerformance(3, 4)
	perm.DisplaySeats()
	go perm.ReserveSeat(1, 1, &mutex)
	go perm.ReserveSeat(1, 1, &mutex)
	time.Sleep(5 * time.Second)
	perm.DisplaySeats()
	time.Sleep(5 * time.Second)
}
