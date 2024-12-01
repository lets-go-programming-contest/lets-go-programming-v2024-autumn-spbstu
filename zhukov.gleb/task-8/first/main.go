package main

func plus(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func main() {
	a, b := 5, 10

	println(plus(a, b))
	println(minus(a, b))
}
