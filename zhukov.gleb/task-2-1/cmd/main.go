package main

import (
	"fmt"
	"os"
	"task-2-1/internal/control"
	myReader "task-2-1/internal/reader"
)

func main() {
	r := myReader.NewConsoleReader()
	controller := control.NewConsoleControl(*r)

	if err := controller.Run(); err != nil {
		fmt.Printf("Произошла ошибка: %v\n", err)
		os.Exit(1)
	}
}
