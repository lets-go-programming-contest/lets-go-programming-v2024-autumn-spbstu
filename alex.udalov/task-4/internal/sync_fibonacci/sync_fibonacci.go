package syncfibonacci

import (
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
	for i := range m.data {
		m.lock.Lock()
		n := rand.Intn(100)
		m.data[i] = append(m.data[i], n)
		m.lock.Unlock()
	}
}

func (m *Matrix) GetMatrix() [][]int {
	return m.data
}
