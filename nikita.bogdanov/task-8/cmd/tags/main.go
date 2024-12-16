package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("b mustn't be zero")
	}
	return a / b
}

func main() {
	var (
		a = 50
		b = 2
	)
	fmt.Println(divide(a, b))
}
