//go:build pro
// +build pro

package main

import "fmt"

func init() {
	fmt.Printf("a = %v, b = %v\n", a, b)
}
