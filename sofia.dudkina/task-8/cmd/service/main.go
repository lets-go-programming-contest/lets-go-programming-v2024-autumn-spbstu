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

func main() {
	var result = add(1, 2)
	println(result)
	result = sub(6, 2)
	println(result)
	result = div(20, 2)
	println(result)
	result = mul(12, 2)
	println(result)
}
