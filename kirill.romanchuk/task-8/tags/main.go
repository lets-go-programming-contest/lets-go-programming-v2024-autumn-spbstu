package main

//go build -o calculator.exe -tags en
import (
	"fmt"
)

func readNum(message string) float64 {
	var num float64
	for {
		fmt.Print(message)
		_, err := fmt.Scan(&num)
		if err == nil {
			break
		}
		fmt.Println(getErrorMessage())
		var dummy string
		fmt.Scanln(&dummy) // Очищаем ввод
	}
	return num
}

func readOperator() string {
	var operator string
	for {
		fmt.Print("Выберите операцию (+, -, *, /): ")
		_, err := fmt.Scan(&operator)
		if err == nil && isOperator(operator) {
			break
		}
		fmt.Println(getOperatorErrorMessage())
		var dummy string
		fmt.Scanln(&dummy) // Очищаем ввод
	}
	return operator
}

func isOperator(operator string) bool {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	return operators[operator]
}

func calculate(num1 float64, num2 float64, operator string) (float64, error) {
	operations := map[string]func(float64, float64) (float64, error){
		"+": func(a, b float64) (float64, error) { return a + b, nil },
		"-": func(a, b float64) (float64, error) { return a - b, nil },
		"*": func(a, b float64) (float64, error) { return a * b, nil },
		"/": func(a, b float64) (float64, error) {
			if b == 0 {
				return 0, fmt.Errorf(getDivideByZeroError())
			}
			return a / b, nil
		},
	}

	if opFunc, exists := operations[operator]; exists {
		return opFunc(num1, num2)
	}
	return 0, fmt.Errorf(getInvalidOperatorError(operator))
}

func readCommand() string {
	var cmd string
	for {
		fmt.Print(getCommandMessage())
		fmt.Scan(&cmd)
		return cmd
	}
}

func main() {
	for {
		cmd := readCommand()
		switch cmd {
		case "0":
			fmt.Println(getFarewellMessage())
			return
		case "1":
			num1 := readNum(getInputMessage("Введите первое число: "))
			operator := readOperator()
			num2 := readNum(getInputMessage("Введите второе число: "))
			result, err := calculate(num1, num2, operator)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Результат: %.3f %s %.3f = %.3f\n", num1, operator, num2, result)
		default:
			fmt.Println(getInvalidCommandMessage())
		}
	}
}
