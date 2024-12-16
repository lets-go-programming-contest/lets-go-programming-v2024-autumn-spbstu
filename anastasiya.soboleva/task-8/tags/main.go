package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number>")
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num < 0 {
		fmt.Println("Please provide a valid positive integer.")
		return
	}

	fmt.Printf("Fibonacci(%d) = %d\n", num, fibonacci(num))
}
