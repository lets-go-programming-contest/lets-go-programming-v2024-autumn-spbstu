package main

import (
	"fmt"
	"os"
	"task-2-2/internal/lunch"
	"task-2-2/internal/reader"
)

func main() {
	r := reader.NewConsoleReader()
	lunchRunner := lunch.NewConsoleLunch(*r)

	if err := lunchRunner.Run(); err != nil {
		fmt.Printf("Произошла ошибка: %v\n", err)
		os.Exit(1)
	}
}
