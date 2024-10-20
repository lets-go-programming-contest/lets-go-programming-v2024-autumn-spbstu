package main

import (
	"fmt"
	"os"

	"github.com/Piyavva/task-2-1/internal/input"
	"github.com/Piyavva/task-2-1/internal/temp"
)

func main() {
    n, err := input.ReadInt(os.Stdin)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    err = temp.GetTemp(os.Stdin, os.Stdout, n)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}