package main

import (
    "fmt"
    "os"
)

func getExpression() (first, second float32, sign string) {
    fmt.Print("Введите первое число: ")
    _, err := fmt.Scan(&first)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Некорректное число. Пожалуйста, введите числовое значение.")
        os.Exit(1)
    }
    fmt.Print("Введите операцию (+, -, *, /): ")
    _, err = fmt.Scan(&sign)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Некорректный ввод. Пожалуйста, введите корректный знак.")
        os.Exit(1)
    }
    switch (sign) {
    case "+":
    case "-":
    case "*":
    case "/":
    default:
        fmt.Fprintln(os.Stderr, "Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
        os.Exit(1)
    }

    fmt.Print("Введите второе число: ")
    _, err = fmt.Scan(&second)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Некорректное число. Пожалуйста, введите числовое значение.")
        os.Exit(1)
    }
    return 
}

func calculate(first, second float32, sign string) (result float32) {
    if sign == "/" && second == 0 {
        fmt.Fprintln(os.Stderr, "Ошибка: Деление на нуль.")
        os.Exit(1);
    }
    switch (sign) {
    case "+":
        result = first + second
    case "-":
        result = first - second
    case "*":
        result = first * second
    case "/":
        result = first / second
    default:
        panic("It cant be...")
    }
    return
}

func main() {
    first, second, sign := getExpression()
    result := calculate(first, second, sign)
    fmt.Printf("%g %s %g = %g\n", first, sign, second, result)
}
