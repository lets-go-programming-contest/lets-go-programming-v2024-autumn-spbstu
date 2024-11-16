package unsyncfibonacci

import (
	"math/rand"
	"time"
)

var matrix [][]int

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

func WriteToMatrix(row int, ch chan int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(20)
	fib := Fibonacci(n)

	if row%2 == 0 {
		time.Sleep(100 * time.Millisecond)
		matrix[row] = append(matrix[row], fib)
		ch <- fib
	} else {
		time.Sleep(100 * time.Millisecond)
		ch <- fib
		matrix[row] = append(matrix[row], fib)

		if row == 2 {
			select {}
		}
	}
}

func GetMatrix() [][]int {
	return matrix
}
