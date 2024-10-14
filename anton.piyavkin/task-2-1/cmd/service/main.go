package main

import (
	"fmt"
	"os"
)

func main() {
    n, err := input.readInt(os.Stdin)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
    err = temp.getTemp(os.Stdin, os.Stdout, n)
    if err != nil {
        fmt.Fprint(os.Stderr, err)
        os.Exit(1)
    }
}