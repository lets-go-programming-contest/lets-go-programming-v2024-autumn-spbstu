package performance

import (
	"fmt"
)

type Performance struct {
	Seats          [][]int
	availableSeats int
}

func CreatePerformance(rows, cols int) Performance {
	seats := make([][]int, rows)
	for i := range seats {
		seats[i] = make([]int, cols)
		for j := range seats[i] {
			seats[i][j] = 0
		}
	}
	return Performance{Seats: seats, availableSeats: rows * cols}
}

func (p Performance) DisplaySeats() {
	for _, row := range p.Seats {
		fmt.Println(row)
	}
}
