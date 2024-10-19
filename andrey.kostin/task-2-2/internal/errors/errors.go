package errors

import "errors"

var (
	ErrorIncorectInt        = errors.New("> Некорректное значение. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectIntBounds  = errors.New("> Некорректное количество. Значение должно быть в диапазоне от 1 до 10000!")
	ErrorIncorectHeap       = errors.New("> Некорректное значение элемента в куче. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectHeapBounds = errors.New("> Некорректное значение элемента в куче. Значение должно быть в диапазоне от -10000 до 10000!")
)
