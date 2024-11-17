package unsyncfibonacci

import (
	"fmt"
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

func (m *Matrix) FillRandom() {
	for i := range m.data {
		n := rand.Intn(100)
		m.data[i] = append(m.data[i], n)

		if i == 2 {
			select {}
		}
	}
}

func (m *Matrix) GetMatrix() [][]int {
	return m.data
}

func RunUnsyncFibonacci(rows int) {
	var matrix Matrix
	matrix.Init(rows)

	matrix.FillRandom()

	fmt.Println("Результат матрицы (без синхронизации):")
	for _, row := range matrix.GetMatrix() {
		fmt.Println(row)
	}
}
