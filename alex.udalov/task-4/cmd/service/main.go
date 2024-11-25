package main

import (
	"fmt"
	"os"

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
		syncfibonacci.RunSyncFibonacci(rows)

	} else if mode == "unsync" {
		fmt.Println("Запуск несинхронизированного режима:")
		unsyncfibonacci.RunUnsyncFibonacci(rows)

	} else {
		fmt.Println("Неизвестный режим. Используйте 'sync' или 'unsync'.")
	}
}
