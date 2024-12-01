package main

var a int

func sum(a int, b int) int {
	return a + b
}

func main() {
	if a != 10 {
		a = 100
	}
	c := sum(a, 5)
	println(c)
}
