//go:build pro
// +build pro

package calculator

func Calculator(number1 float64, number2 float64, operation string) float64 {
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
