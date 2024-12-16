//go:build pro
// +build pro

package main

import (
	"fmt"
)

func init() {
	fmt.Println("Pro version activated!")
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
