package errors

import "errors"

var (
	ErrIncorrectNumber       = errors.New("Некорректное число. Пожалуйста, введите числовое значение.")
	ErrIncorrectSeparator    = errors.New("Строка должна содержать ровно один пробел для разделения.")
	ErrIncorrectComparsionOp = errors.New("Некорректная операция сравнения. Пожалуйста, введите \">=\" или \"<=\".")
)
