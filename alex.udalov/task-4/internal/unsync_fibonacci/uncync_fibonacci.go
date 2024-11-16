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

// Запись случайного числа Фибоначчи в строку матрицы (может вызвать дедлок)
func WriteToMatrix(row int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(20) // Генерируем случайное число от 0 до 19
	fib := Fibonacci(n)

	// Имитация дедлока: две горутины пытаются записать в одну и ту же строку
	if row%2 == 0 { // Если четный номер строки
		time.Sleep(100 * time.Millisecond) // Задержка для имитации конфликта
		matrix[row] = append(matrix[row], fib)
	} else { // Если нечетный номер строки
		time.Sleep(100 * time.Millisecond) // Задержка для имитации конфликта
		matrix[row] = append(matrix[row], fib)
	}

	// Принудительное создание дедлока:
	if row == 2 { // Условие для создания дедлока
		select {} // Бесконечный блокирующий вызов (дедлок)
	}
}

// Получение матрицы (может вернуть некорректные значения!)
func GetMatrix() [][]int {
	return matrix
}
