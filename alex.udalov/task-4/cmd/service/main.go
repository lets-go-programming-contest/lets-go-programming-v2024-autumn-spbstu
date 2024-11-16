package main

import (
	"fmt"
	"time"

	"sync"
	syncfibonacci "task-4/internal/sync_fibonacci"
	unsyncfibonacci "task-4/internal/unsync_fibonacci"
)

func main() {
	const rows = 5 // Количество строк в матрице

	fmt.Println("Вычисление чисел Фибоначчи с использованием Mutex и каналов:")
	syncfibonacci.InitMatrix(rows)

	resultChan := make(chan int, rows) // Канал для передачи результатов

	var wg sync.WaitGroup

	for i := 0; i < rows; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			syncfibonacci.WriteToMatrix(row, resultChan) // Теперь передаем канал в функцию
		}(i)
	}

	wg.Wait()         // Ждем завершения всех горутин
	close(resultChan) // Закрываем канал после завершения всех горутин

	fmt.Println("Результат матрицы (синхронизировано):")
	for result := range resultChan { // Читаем результаты из канала
		fmt.Println(result)
	}
	fmt.Println(syncfibonacci.GetMatrix())

	fmt.Println("\nВычисление чисел Фибоначчи без использования синхронизации (может вызвать дедлок):")
	unsyncfibonacci.InitMatrix(rows)

	var wgUnsync sync.WaitGroup

	for i := 0; i < rows; i++ {
		wgUnsync.Add(1)
		go func(row int) {
			defer wgUnsync.Done()
			time.Sleep(100 * time.Millisecond) // Задержка для демонстрации проблемы с дедлоком
			if row%2 == 0 {                    // Условие для имитации дедлока (попытка записи в одну и ту же строку)
				time.Sleep(200 * time.Millisecond)
			}
			unsyncfibonacci.WriteToMatrix(row) // Здесь произойдет дедлок или зависание!
		}(i)
	}

	wgUnsync.Wait() // Ждем завершения всех горутин

	fmt.Println("Результат матрицы (без синхронизации):")
	for _, row := range unsyncfibonacci.GetMatrix() {
		fmt.Println(row)
	}
}
