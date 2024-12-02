//go:build ru
// +build ru

package main

func getCommandMessage() string {
	return "Введите команду (1 для вычисления, 0 для выхода): "
}

func getInputMessage(field string) string {
	return "Введите " + field + ": "
}

func getFarewellMessage() string {
	return "Выход из программы."
}

func getErrorMessage() string {
	return "Некорректное число. Пожалуйста, введите числовое значение."
}

func getOperatorErrorMessage() string {
	return "Некорректная операция. Пожалуйста, используйте символы +, -, * или /."
}

func getInvalidCommandMessage() string {
	return "Неверная команда. Пожалуйста, попробуйте снова."
}

func getInvalidOperatorError(operator string) string {
	return "Ошибка: недопустимый оператор '" + operator + "'"
}

func getDivideByZeroError() string {
	return "Ошибка: деление на 0"
}
