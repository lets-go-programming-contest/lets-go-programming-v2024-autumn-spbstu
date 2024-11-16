package syncfibonacci

import (
	"math/rand"
	"sync"
	"time"
)

// Глобальные переменные для хранения результатов и мьютекса
var (
	matrix [][]int
	lock   sync.Mutex
)

// Инициализация матрицы заданного размера
func InitMatrix(rows int) {
	matrix = make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, 0) // Инициализируем пустые строки
	}
}

// Функция для вычисления n-го числа Фибоначчи
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

// Запись случайного числа Фибоначчи в строку матрицы через канал
func WriteToMatrix(row int, resultChan chan<- int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(20) // Генерируем случайное число от 0 до 19
	fib := Fibonacci(n)

	lock.Lock() // Блокируем доступ к матрице
	matrix[row] = append(matrix[row], fib)
	lock.Unlock() // Разблокируем доступ к матрице

	resultChan <- fib // Отправляем результат в канал
}

// Получение матрицы
func GetMatrix() [][]int {
	return matrix
}
