package main

import (
	"fmt"
	"os"
	"sync"

	syncfibonacci "task-4/internal/sync_fibonacci"
	unsyncfibonacci "task-4/internal/unsync_fibonacci"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [sync|unsync]")
		return
	}

	mode := os.Args[1]
	const rows = 100

	if mode == "sync" {
		fmt.Println("Запуск синхронизированного режима:")

		var matrix syncfibonacci.Matrix
		matrix.Init(rows)

		var wg sync.WaitGroup

		for i := 0; i < rows; i++ {
			wg.Add(1)
			go func(row int) {
				defer wg.Done()
				matrix.FillRandom()
			}(i)
		}

		wg.Wait()

		fmt.Println("Результат матрицы (синхронизировано):")
		for _, row := range matrix.GetMatrix() {
			fmt.Println(row)
		}

	} else if mode == "unsync" {
		fmt.Println("Запуск несинхронизированного режима:")

		var matrix unsyncfibonacci.Matrix
		matrix.Init(rows)

		ch := make(chan int)

		var wg sync.WaitGroup

		wg.Add(1)
		go func() {
			defer wg.Done()
			matrix.FillRandom(ch)
		}()

		wg.Wait()

		fmt.Println("Результат матрицы (без синхронизации):")
		for _, row := range matrix.GetMatrix() {
			fmt.Println(row)
		}
	} else {
		fmt.Println("Неизвестный режим. Используйте 'sync' или 'unsync'.")
	}
}
