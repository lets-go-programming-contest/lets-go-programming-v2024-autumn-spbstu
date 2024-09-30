package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	calculator := NewCalculator()

	calculator.Run()
}

type Calculator struct {
	reader *bufio.Reader
}

// NewCalculator создает новый экземпляр Calculator
func NewCalculator() *Calculator {
	return &Calculator{reader: bufio.NewReader(os.Stdin)}
}

// readFloat считывает число с плавающей запятой
func (c *Calculator) readFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := c.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}

		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение.")
	}
}

// isOperator проверяет, является ли строка оператором
func (c *Calculator) isOperator(str string) bool {
	operators := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}
	return operators[str]
}

// readOperator считывает арифметический оператор
func (c *Calculator) readOperator(prompt string) string {
	for {
		fmt.Print(prompt)
		input, _ := c.reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if c.isOperator(input) {
			return input
		}

		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
}

// Run запускает калькулятор
func (c *Calculator) Run() {
	for {
		// Читаем первое число
		operand1 := c.readFloat("Введите первое число: ")

		// Читаем оператор
		operator := c.readOperator("Выберите операцию (+, -, *, /): ")

		// Читаем второе число
		operand2 := c.readFloat("Введите второе число: ")

		// Проверка на деление на ноль
		if operator == "/" && operand2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно.")
			continue
		}

		// Выполняем операцию
		var result float64
		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			result = operand1 / operand2
		}

		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n\n", operand1, operator, operand2, result)
	}
}
