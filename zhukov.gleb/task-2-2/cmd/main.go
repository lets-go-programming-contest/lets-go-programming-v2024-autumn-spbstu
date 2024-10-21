package main

import (
	"task-2-2/internal/lunch"
)

func main() {
	lunchRunner := lunch.NewConsoleLunch()
	lunchRunner.Run()
}
