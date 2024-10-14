package errors

import "errors"

var (
	ErrorIncorectInt       = errors.New("> Некорректное значение. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectIntBounds = errors.New("> Некорректное количество. Значение должно быть в диапазоне от 1 до 1000!")
	ErrorIncorectCondComp  = errors.New("> Некорректное условие. Ожидалось '<=' или '>='.")
	ErrorIncorectCondSpace = errors.New("> Некорректный формат. Убедитесь, что ввели условие из двух частей через пробел.")
	ErrorIncorectCondTemp  = errors.New("> Некорректное значение температуры. Ожидалось целочисленное значение.")
)
