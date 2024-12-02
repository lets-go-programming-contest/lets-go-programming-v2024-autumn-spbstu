package main

import (
	"fmt"

	"github.com/solomonalfred/task-3/internal/application"
)

func main() {
	err := application.App()
	if err != nil {
		fmt.Println(err)
	}
}
