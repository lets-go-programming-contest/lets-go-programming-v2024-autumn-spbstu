package main

import (
	"task-3/internal/controller"
	"task-3/internal/fileWorkers"
)

func main() {
	p := fileWorkers.NewParser()
	control := controller.NewController(*p)

	if err := control.Run(); err != nil {
		panic(err)
	}
}
