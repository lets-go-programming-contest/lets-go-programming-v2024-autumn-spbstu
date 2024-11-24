package performance

import (
	"fmt"
	"sync"
	"time"
)

type Performance struct {
	seats [][]int
	Rows  int
	Cols  int
}

func CreatePerformance(rows, cols int) Performance {
	seats := make([][]int, rows)
	for i := range seats {
		seats[i] = make([]int, cols)
		for j := range seats[i] {
			seats[i][j] = 0
		}
	}
	return Performance{seats: seats, Rows: rows, Cols: cols}
}

func (p Performance) DisplaySeats() {
	for _, row := range p.seats {
		fmt.Println(row)
	}
}

func (p *Performance) ReserveSeat(row, col int, mutex *sync.Mutex, result chan<- bool) {
	//mutex.Lock()
	if p.seats[row][col] == 0 {
		time.Sleep(1 * time.Second) // Simulation of payment
		p.seats[row][col] += 1
		result <- true
	} else {
		result <- false
	}
	//mutex.Unlock()
}
