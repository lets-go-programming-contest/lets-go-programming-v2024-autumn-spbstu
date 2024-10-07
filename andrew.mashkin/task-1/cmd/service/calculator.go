package main

func Calcualte(first int, second int, operation string) int {
	switch operation {
	case "+":
		return first + second
	case "-":
		return first - second
	case "*":
		return first * second
	case "/":
		return first / second
	default:
		return 0
	}
}
