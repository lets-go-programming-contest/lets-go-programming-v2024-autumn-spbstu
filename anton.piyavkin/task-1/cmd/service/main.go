package main

import (
	"errors"
	"fmt"
	"os"
)

func calculate(lhs int, rhs int, op string) (int, error) {
    switch op {
    case "+":
        return lhs + rhs, nil
    case "-":
        return lhs - rhs, nil
    case "/":
        if rhs == 0 {
            return 0, errors.New("деление на 0")
        }
        return lhs / rhs, nil
    default:
        return lhs * rhs, nil
    }
}

func main() {
    fmt.Print("Введите первое число:")
    var lhs int
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
    var rhs int
    fmt.Print("Введите второе число:")
    _, err = fmt.Scan(&rhs)
    if err != nil {
    	fmt.Fprint(os.Stderr, "Некорректное число. Пожалуйста, введите числовое значение.\n")
    	return
    }
    var res int
    res, err = calculate(lhs, rhs, op)
    if err != nil {
        fmt.Fprint(os.Stderr, err.Error())
        return
    }
    fmt.Println("Результат:", lhs, op, rhs, "=", res)
}