package internal

import "errors"

var Operations = map[string]func(float64, float64) (float64, error){
	"+": add,
	"-": subtract,
	"*": multiply,
	"/": divide,
}

func add(a, b float64) (float64, error) {
	return a + b, nil
}

func subtract(a, b float64) (float64, error) {
	return a - b, nil
}

func multiply(a, b float64) (float64, error) {
	return a * b, nil
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль невозможно")
	}
	return a / b, nil
}
