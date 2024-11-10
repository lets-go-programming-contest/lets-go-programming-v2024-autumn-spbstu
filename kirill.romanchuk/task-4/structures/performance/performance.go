package performance

import "fmt"

type Performance struct {
	Seats [][]bool
}

func CreatePerformance(rows, cols int) Performance {
	seats := make([][]bool, rows)
	for i := range seats {
		seats[i] = make([]bool, cols)
		for j := range seats[i] {
			seats[i][j] = true
		}
	}
	return Performance{Seats: seats}
}

func (p Performance) DisplaySeats() {
	for _, row := range p.Seats {
		fmt.Println(row)
	}
}
