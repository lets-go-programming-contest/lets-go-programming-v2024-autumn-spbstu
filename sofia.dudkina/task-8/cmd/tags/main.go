package main

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func div(a, b float64) float64 {
	return a / b
}

func mul(a, b float64) float64 {
	return a * b
}

var result = 1.0

func main() {
	result += sub(6, 2)
	result += mul(6, 2)
	result -= sub(16, 2)
	println(result)
}
