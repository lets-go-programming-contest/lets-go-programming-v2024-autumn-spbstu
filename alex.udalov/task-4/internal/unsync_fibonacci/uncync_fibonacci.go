package syncfibonacci

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

func fibonacci(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func (m *Matrix) FillRandom() {

	for i := range m.data {
		go func(row int) {
			rowData := make([]int, 10000000)
			for j := 0; j < 10000000; j++ {
				n := rand.Intn(20)
				rowData[j] = fibonacci(n)
			}

			m.data[row] = rowData
		}(i)
	}

}

func (m *Matrix) GetMatrix() [][]int {
	return m.data
}

func RunUnsyncFibonacci(rows int) {
	var matrix Matrix
	matrix.Init(rows)

	matrix.FillRandom()

	success := true
	for _, row := range matrix.GetMatrix() {
		if len(row) != 10000 {
			success = false
			break
		}
	}
	if success {
		fmt.Println("Результат матрицы (без синхронизации): true")
	} else {
		fmt.Println("Результат матрицы (без синхронизации): false")
	}
}
