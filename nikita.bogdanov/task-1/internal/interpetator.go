package internal

import (
	"math/big"
)

const (
	INT = iota
	FLOAT
	COMPLEX
)

func interpret(num1, num2 string) int {
	var (
		type_num1 = checkType(num1)
		type_num2 = checkType(num2)
	)

	var res_type int
	if type_num1 >= type_num2 {
		res_type = type_num1
	} else {
		res_type = type_num2
	}
	return res_type
}

func checkType(num string) int {
	var intValue big.Int
	_, ok_int := intValue.SetString(num, 10)
	if ok_int {
		return INT
	}

	var floatValue big.Float
	_, ok_float := floatValue.SetString(num)
	if ok_float {
		return FLOAT
	}

	_, ok_complex := ParseComplex(num)
	if ok_complex {
		return COMPLEX
	}

	panic("Некорректное число. Пожалуйста, введите числовое значение.")
}
