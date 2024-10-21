package main

import (
	"task-2-1/internal/control"
)

func main() {
	controller := control.NewConsoleControl()
	controller.Run()
}
