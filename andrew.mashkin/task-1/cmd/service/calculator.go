package main

import (
	"errors"
	"fmt"
)

func Calcualte(first int, second int, operation string) (int, error) {
	switch operation {
	case "+":
		return Add(first, second)
	case "-":
		return Subtitle(first, second)
	case "*":
		return Multiply(first, second)
	case "/":
		return Divide(first, second)
	default:
		return 0, errors.New("Некорректная операция. Пожалуйста, используйте символы +, -, * или /")
	}
}

func Add(first int, second int) (int, error) {
	if (second > 0 && first > (1<<31-1)-second) || (second < 0 && first < (-1<<31)-second) {
		return 0, errors.New("возникло переполнение хыхы")
	}
	return first + second, nil
}

func Subtitle(first, second int) (int, error) {
	if (second > 0 && first < (-1<<31)+second) || (second < 0 && first > (1<<31-1)-second) {
		return 0, errors.New("возникло переполнение хыхы")
	}
	return first - second, nil
}

func Multiply(first, second int) (int, error) {
	if first > 0 && second > 0 && first > (1<<31-1)/second {
		return 0, errors.New("возникло переполнение хыхы")
	}
	if first < 0 && second < 0 && first < (-1<<31)/second {
		return 0, errors.New("возникло переполнение хыхы")
	}
	if (first > 0 && second < 0 && first > (-1<<31)/second) || (first < 0 && second > 0 && first < (1<<31-1)/second) {
		return 0, errors.New("возникло переполнение хыхы")
	}
	return first * second, nil
}

func Divide(first, second int) (int, error) {
	if second == 0 {
		return 0, fmt.Errorf("ошибка: деление на ноль")
	}
	return first / second, nil
}
