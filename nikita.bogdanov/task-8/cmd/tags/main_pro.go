//go:build pro
// +build pro

package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func init() {
	var (
		a = 50
		b = 2
	)
	fmt.Println(add(a, b))
}
