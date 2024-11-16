package syncfibonacci

import (
	"math/rand"
	"sync"
	"time"
)

var (
	matrix [][]int
	lock   sync.Mutex
)

func InitMatrix(rows int) {
	matrix = make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, 0)
	}
}

func Fibonacci(n int) int {
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

func WriteToMatrix(row int, resultChan chan<- int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(20)
	fib := Fibonacci(n)

	lock.Lock()
	matrix[row] = append(matrix[row], fib)
	lock.Unlock()

	resultChan <- fib
}

func GetMatrix() [][]int {
	return matrix
}
