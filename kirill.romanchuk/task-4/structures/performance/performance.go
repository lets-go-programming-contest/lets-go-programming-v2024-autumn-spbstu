package performance

import (
	"fmt"
	"sync"
	"time"
)

type Performance struct {
	Seats [][]int
}

func CreatePerformance(rows, cols int) Performance {
	seats := make([][]int, rows)
	for i := range seats {
		seats[i] = make([]int, cols)
		for j := range seats[i] {
			seats[i][j] = 0
		}
	}
	return Performance{Seats: seats}
}

func (p Performance) DisplaySeats() {
	for _, row := range p.Seats {
		fmt.Println(row)
	}
}

func (p *Performance) ReserveSeat(row, col int, mutex *sync.Mutex) {
	mutex.Lock()
	if p.Seats[row][col] == 0 {
		time.Sleep(2 * time.Second) // Simulated payment
		p.Seats[row][col] += 1
	}
	mutex.Unlock()
}
