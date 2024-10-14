package internal

import (
	"fmt"
	"math/big"
)

const (
	FLOAT_PREC = 1000
)

func CalcExpression(num1, num2, operator string) {
	num_type := interpret(num1, num2)
	switch num_type {
	case INT:
		number1 := new(big.Int)
		number1.SetString(num1, 10)
		number2 := new(big.Int)
		number2.SetString(num2, 10)
		IntCalc(number1, number2, operator)
	case FLOAT:
		number1 := new(big.Float).SetPrec(FLOAT_PREC)
		number1.SetString(num1)
		number2 := new(big.Float).SetPrec(FLOAT_PREC)
		number2.SetString(num2)
		FloatCalc(number1, number2, operator)
	case COMPLEX:
		number1, _ := ParseComplex(num1)
		number2, _ := ParseComplex(num2)
		ComplexCalc(number1, number2, operator)
	}
}

func IntCalc(num1, num2 *big.Int, operator string) {
	result := new(big.Int)
	switch operator {
	case "+":
		result.Add(num1, num2)
	case "-":
		result.Sub(num1, num2)
	case "*":
		result.Mul(num1, num2)
	case "/":
		if num2.Cmp(big.NewInt(0)) == 0 {
			panic("Ошибка: деление на ноль невозможно")
		}
		result.Div(num1, num2)
	default:
		panic("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
	fmt.Print(result.Text(10))
}

func FloatCalc(num1, num2 *big.Float, operator string) {
	result := new(big.Float)
	switch operator {
	case "+":
		result.Add(num1, num2)
	case "-":
		result.Sub(num1, num2)
	case "*":
		result.Mul(num1, num2)
	case "/":
		if num2.Cmp(big.NewFloat(0)) == 0 {
			panic("Ошибка: деление на ноль невозможно")
		}
		result.Quo(num1, num2)
	default:
		panic("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
	fmt.Print(result.Text('f', -1))
}

func ComplexCalc(num1, num2 *BigComplex, operator string) {
	result := num1
	switch operator {
	case "+":
		result = num1.Add(num2)
	case "-":
		result = num1.Sub(num2)
	case "*":
		result = num1.Mul(num2)
	case "/":
		result = num1.Div(num2)
	default:
		panic("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
	}
	fmt.Print(result)
}
