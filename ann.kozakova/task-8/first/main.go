package main

func main() {

	var res float64
	res = operate(3, 2, "+")
	println(res)
	res = operate(4, 5, "-")
	println(res)
	res = operate(5, 6, "")
	println(res)
}

func operate(a, b float64, operator string) float64 {
	switch operator {
	case "+":
		return plus(a, b)
	case "-":
		return minus(a, b)
	default:
		return multiply(a, b)
	}
}

func plus(a, b float64) float64 {
	return a + b
}

func minus(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}
