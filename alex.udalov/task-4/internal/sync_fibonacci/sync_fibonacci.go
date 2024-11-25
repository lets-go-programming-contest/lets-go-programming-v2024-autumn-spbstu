package syncfibonacci

import (
	"fmt"
	"math/rand"
	"sync"
)

type Matrix struct {
	data [][]int
	lock sync.Mutex
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
	var wg sync.WaitGroup

	for i := range m.data {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()

			rowData := make([]int, 5)
			for j := 0; j < 5; j++ {
				n := rand.Intn(20)
				rowData[j] = fibonacci(n)
			}

			m.lock.Lock()
			m.data[row] = rowData
			m.lock.Unlock()
		}(i)
	}

	wg.Wait()
}

func (m *Matrix) GetMatrix() [][]int {
	return m.data
}

func RunSyncFibonacci(rows int) {
	var matrix Matrix
	matrix.Init(rows)

	matrix.FillRandom()

	fmt.Println("Результат матрицы (синхронизировано):")
	for _, row := range matrix.GetMatrix() {
		fmt.Println(row)
	}
}
