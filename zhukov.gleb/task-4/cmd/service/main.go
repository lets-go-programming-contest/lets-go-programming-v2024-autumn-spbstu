package main

import (
	"fmt"
	"task-4/pkg/metro"
)

func main() {
	textChan := make(chan string)

	vRepo := metro.NewVisitorsRepo()
	go vRepo.Simulator(100, textChan)

	for msg := range textChan {
		fmt.Println(msg)
	}
}
