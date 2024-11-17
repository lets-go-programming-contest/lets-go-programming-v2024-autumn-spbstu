package unsyncfibonacci

import (
	"math/rand"
)

type Matrix struct {
	data [][]int
}

func (m *Matrix) Init(rows int) {
	m.data = make([][]int, rows)
	for i := range m.data {
		m.data[i] = make([]int, 0)
	}
}

func (m *Matrix) FillRandom(ch chan int) {
	for i := range m.data {
		n := rand.Intn(100)
		m.data[i] = append(m.data[i], n)

		ch <- n
	}
}

func (m *Matrix) GetMatrix() [][]int {
	return m.data
}
