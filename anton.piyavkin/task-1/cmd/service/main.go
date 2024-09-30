package main

import (
	"fmt"
	"os"
)

func main() {
    fmt.Print("Введите первое число:")
	var lhs int64
	_, err := fmt.Scan(&lhs)
	if err != nil {
		fmt.Fprint(os.Stderr, "Некорректное число. Пожалуйста, введите числовое значение.\n")
		return
	}
	fmt.Print("Выберите операцию (+, -, *, /):")
	var op string
	_, err = fmt.Scan(&op)
	if err != nil || (op != "+" && op != "*" && op != "-" && op != "/") {
		fmt.Fprint(os.Stderr, "Некорректная операция. Пожалуйста, используйте символы +, -, * или /.\n")
		return
	}
	var rhs int64
	fmt.Print("Введите второе число:")
	_, err = fmt.Scan(&rhs)
	if err != nil || (op == "/" && rhs == 0) {
		fmt.Fprint(os.Stderr, "Некорректное число. Пожалуйста, введите числовое значение.\n")
		return
	}
}