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
		if input == "+" || input == "*" || input == "/" || input == "-" {
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
		firstOperand, _ := strconv.ParseFloat(userInputs[0], 64)
		operator := userInputs[1]
		secondOperand, _ := strconv.ParseFloat(userInputs[2], 64)

		var result float64
		switch operator {
		case "+":
			result = firstOperand + secondOperand
		case "-":
			result = firstOperand - secondOperand
		case "*":
			result = firstOperand * secondOperand
		case "/":
			if secondOperand == 0 {
				fmt.Println("+++++Ошибка: Деление на ноль+++++")
				return
			}
			result = firstOperand / secondOperand
		default:
			fmt.Println("+++++Ошибка: Неподдерживаемый оператор+++++")
			return
		}

		fmt.Printf("Результат: %.2f\n", result)
	}

	var textForUser = map[int]string{
		0: "Введите первый операнд:",
		1: "Введите оператор:",
		2: "Введите второй операнд:",
	}

LOOP:
	for {
		userInputs := make(map[int]string, 3)
		for i := 0; i < 3; i++ {
			fmt.Println(textForUser[i])
			fmt.Print("> ")
			userInputs[i] = reader()
			if endProg(userInputs[i]) {
				break LOOP
			}
			if !inputIsValid(userInputs[i], i%2 == 1) {
				fmt.Println("+++++Ошибка ввода. Вводите корректные данные+++++")
				break
			}
		}
		resultsUserInput(userInputs)
	}
}
