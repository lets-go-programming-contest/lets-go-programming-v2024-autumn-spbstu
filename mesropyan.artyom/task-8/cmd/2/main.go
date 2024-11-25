package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

var a = 2
var b = 3

func main() {
	fmt.Println(sum(a, b))
}
