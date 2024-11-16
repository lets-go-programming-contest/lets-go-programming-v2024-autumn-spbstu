package unsyncfibonacci

import (
	"math/rand"
	"time"
)

// Глобальные переменные для хранения результатов (без мьютекса)
var matrix [][]int

// Инициализация матрицы заданного размера
func InitMatrix(rows int) {
	matrix = make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, 0) // Инициализируем пустые строки
	}
}

// Функция для вычисления n-го числа Фибоначчи без синхронизации (может быть некорректной)
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
	return b // Некорректный результат при многопоточном вызове!
}

// Запись случайного числа Фибоначчи в строку матрицы (может быть некорректной)
func WriteToMatrix(row int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(20) // Генерируем случайное число от 0 до 19
	fib := Fibonacci(n)

	matrix[row] = append(matrix[row], fib) // Некорректный доступ к матрице без мьютекса
}

// Получение матрицы (может вернуть некорректные значения!)
func GetMatrix() [][]int {
	return matrix
}
