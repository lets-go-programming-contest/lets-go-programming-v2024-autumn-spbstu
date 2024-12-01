package main

func calculator(number1 float64, number2 float64, operation string) float64 {
	switch operation {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "*":
		return number1 * number2
	default:
		return number1 / number2
	}
}

func main() {
	var (
		number1   float64 = 10.23
		number2   float64 = 1.9
		operation string  = "/"
	)
	calculator(number1, number2, operation)
}
