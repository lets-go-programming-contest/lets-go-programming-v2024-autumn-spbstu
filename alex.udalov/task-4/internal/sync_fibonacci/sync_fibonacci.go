package syncfibonacci

import (
	"fmt"
	"math/rand"
	"sync"
)

type Matrix struct {
	data [][]int
	lock sync.Mutex
	once sync.Once
}

func (m *Matrix) Init(rows int) {
	m.once.Do(func() {
		m.data = make([][]int, rows)
		for i := range m.data {
			m.data[i] = make([]int, 0)
		}
	})
}

func (m *Matrix) FillRandom() {
	var wg sync.WaitGroup

	for i := range m.data {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				m.lock.Lock()
				n := rand.Intn(100)
				m.data[row] = append(m.data[row], n)
				m.lock.Unlock()
			}
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
