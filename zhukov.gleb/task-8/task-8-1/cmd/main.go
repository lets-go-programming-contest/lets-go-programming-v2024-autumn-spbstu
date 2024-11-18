package main

func dec(a, b int) int {
	return a - b
}

func inc(a, b int) int {
	return a + b
}

func main() {
	a, b := 5, 10

	println(dec(a, b))
	println(inc(a, b))
}
