package main

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
