package main

import (
	"fmt"
	"os"

	"task-2-2/internal/app"
)

func main() {
	in := os.Stdin
	out := os.Stdout
	err := app.Application(in, out)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
