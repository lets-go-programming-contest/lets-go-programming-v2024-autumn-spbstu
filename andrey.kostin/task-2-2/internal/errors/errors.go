package errors

import "errors"

var (
	ErrorIncorectInt            = errors.New("> Некорректное значение количества. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectIntBounds      = errors.New("> Некорректное количество. Значение должно быть в диапазоне от 1 до 10000!")
	ErrorIncorectHeap           = errors.New("> Некорректное значение предпочтения. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectHeapBounds     = errors.New("> Некорректный значение предпочтения. Значение должно быть в диапазоне от -10000 до 10000!")
	ErrorIncorectPrefDish       = errors.New("> Некорректное значение номера. Пожалуйста, введите целочисленное значение!")
	ErrorIncorectPrefDishBounds = errors.New("> Некорректный номер блюда. Значение не должно быть больше количества блюд!")
)
