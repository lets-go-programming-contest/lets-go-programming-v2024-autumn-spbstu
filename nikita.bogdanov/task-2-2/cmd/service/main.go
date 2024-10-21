package main

import (
	"fmt"
	"os"
	app "task-2-2/internal"
)

func main() {
	in := os.Stdin
	out := os.Stdout
	err := app.Application(in, out)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}
