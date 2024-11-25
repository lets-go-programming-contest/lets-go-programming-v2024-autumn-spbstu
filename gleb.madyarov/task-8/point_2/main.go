package main

var operators = []string{
	"+",
	"-",
}

func main() {
	var res int
	for _, op := range operators {
		res = calc(15, 5, op)
		println(res)
	}
}

func calc(value1 int, value2 int, operator string) int {
	switch operator {
	case "+":
		return add(value1, value2)
	case "-":
		return sub(value1, value2)
	case "*":
		return mul(value1, value2)
	case "/":
		return div(value1, value2)
	default:
		return 0
	}

}

func add(value1 int, value2 int) int {
	return value1 + value2
}

func sub(value1 int, value2 int) int {
	return value1 - value2
}

func mul(value1 int, value2 int) int {
	return value1 * value2
}

func div(value1 int, value2 int) int {
	return value1 / value2
}
