package main

import (
	"fmt"

	"task-4/pkg/metro"
	unsafe "task-4/pkg/unsafemetro"
)

func main() {
	var input int
	fmt.Println("1 - safe; 2 - unsafe")
	fmt.Scanln(&input)

	textChan := make(chan string)

	if input == 1 {
		vRepo := metro.NewVisitorsRepo()
		go vRepo.Simulator(100, textChan)
	} else if input == 2 {
		vRepo := unsafe.UnsafeNewVisitorsRepo()
		go vRepo.UnsafeSimulator(10000, textChan)
	} else {
		return
	}

	for msg := range textChan {
		fmt.Println(msg)
	}
}
