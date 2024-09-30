package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор. Вводите в соответствии с запросами программы")
	fmt.Println("'выход' - для завершения")

	var actions = map[string]func(firstOperand, secondOperand float64) float64{
		"*": func(firstOperand, secondOperand float64) float64 {
			return firstOperand * secondOperand
		},
		"+": func(firstOperand, secondOperand float64) float64 {
			return firstOperand + secondOperand
		},
		"-": func(firstOperand, secondOperand float64) float64 {
			return firstOperand - secondOperand
		},
		"/": func(firstOperand, secondOperand float64) float64 {
			return firstOperand / secondOperand
		},
	}

	inputIsValid := func(input string, isOperator bool) bool {
		if !isOperator {
			input = strings.ReplaceAll(input, ",", ".")
			if _, err := strconv.Atoi(input); err == nil {
				return true
			}
			if _, err := strconv.ParseFloat(input, 64); err == nil {
				return true
			}
			return false
		}
		if _, exists := actions[input]; exists {
			return true
		}
		return false
	}

	endProg := func(input string) bool {
		return input == "выход"
	}

	reader := func() string {
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		return strings.TrimSpace(input)
	}

	resultsUserInput := func(userInputs map[int]string) {
		firstOperand, _ := strconv.ParseFloat(strings.ReplaceAll(userInputs[0], ",", "."), 64)
		operator := userInputs[1]
		secondOperand, _ := strconv.ParseFloat(strings.ReplaceAll(userInputs[2], ",", "."), 64)

		if action, exists := actions[operator]; exists {
			if operator == "/" && secondOperand == 0 {
				fmt.Println("+++++Ошибка: Деление на ноль+++++")
				return
			}
			result := action(firstOperand, secondOperand)
			fmt.Printf("Результат: %.2f\n", result)
		} else {
			fmt.Println("+++++Ошибка: Неподдерживаемый оператор+++++")
		}
	}

	var textForUser = map[int]string{
		0: "Введите первый операнд:",
		1: "Введите оператор:",
		2: "Введите второй операнд:",
	}

	for {
		userInputs := make(map[int]string, 3)
		var validInput bool
		for i := 0; i < 3; i++ {
			fmt.Println(textForUser[i])
			fmt.Print("> ")
			userInputs[i] = reader()
			if endProg(userInputs[i]) {
				return
			}
			validInput = inputIsValid(userInputs[i], i%2 == 1)
			if !validInput {
				fmt.Println("+++++Ошибка ввода. Вводите корректные данные+++++")
				break
			}
		}
		if validInput {
			resultsUserInput(userInputs)
		}
	}
}
