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
	const rows = 10

	if mode == "sync" {
		fmt.Println("Запуск синхронизированного режима:")
		syncfibonacci.InitMatrix(rows)

		resultChan := make(chan int, rows)

		var wg sync.WaitGroup

		for i := 0; i < rows; i++ {
			wg.Add(1)
			go func(row int) {
				defer wg.Done()
				syncfibonacci.WriteToMatrix(row, resultChan)
			}(i)
		}

		wg.Wait()
		close(resultChan)

		fmt.Println("Результат матрицы (синхронизировано):")
		for result := range resultChan {
			fmt.Println(result)
		}

	} else if mode == "unsync" {
		fmt.Println("Запуск несинхронизированного режима:")
		unsyncfibonacci.InitMatrix(rows)

		var wgUnsync sync.WaitGroup

		ch := make(chan int)

		for i := 0; i < rows; i++ {
			wgUnsync.Add(1)
			go func(row int) {
				defer wgUnsync.Done()
				unsyncfibonacci.WriteToMatrix(row, ch)
			}(i)
		}

		wgUnsync.Wait()

		fmt.Println("Результат матрицы (без синхронизации):")
		for _, row := range unsyncfibonacci.GetMatrix() {
			fmt.Println(row)
		}
	} else {
		fmt.Println("Неизвестный режим. Используйте 'sync' или 'unsync'.")
	}
}
