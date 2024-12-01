package main

func main() {
	res := 0
	res = calc(10, 12, "+")
	println(res)
	res = calc(1000, 7, "-")
	println(res)
}

func calc(value1 int, value2 int, operator string) int {
	switch operator {
	case "+":
		return add(value1, value2)
	case "-":
		return sub(value1, value2)
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
