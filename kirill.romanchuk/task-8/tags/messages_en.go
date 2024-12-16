//go:build !ru
// +build !ru

package main

func getCommandMessage() string {
	return "Enter command (1 to calculate, 0 to exit): "
}

func getInputMessage(field string) string {
	return "Enter " + field + ": "
}

func getFarewellMessage() string {
	return "Exiting the program."
}

func getErrorMessage() string {
	return "Invalid number. Please enter a numeric value."
}

func getOperatorErrorMessage() string {
	return "Invalid operation. Please use +, -, * or /."
}

func getInvalidCommandMessage() string {
	return "Invalid command. Please try again."
}

func getInvalidOperatorError(operator string) string {
	return "Error: invalid operator '" + operator + "'"
}

func getDivideByZeroError() string {
	return "Error: division by zero"
}
