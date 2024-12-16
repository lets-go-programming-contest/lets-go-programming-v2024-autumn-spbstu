package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {

	if len(os.Args) != 3 {
		log.Fatal("Usage: go run main.go <temperature> <unit>\n<Unit>: 'C' for Celsius, 'F' for Fahrenheit")
	}

	temperature, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Invalid temperature input: %v\n", err)
	}

	unit := os.Args[2]

	switch unit {
	case "C":
		fahrenheit := celsiusToFahrenheit(temperature)
		fmt.Printf("%.2f°C = %.2f°F\n", temperature, fahrenheit)
	case "F":
		celsius := fahrenheitToCelsius(temperature)
		fmt.Printf("%.2f°F = %.2f°C\n", temperature, celsius)
	default:
		log.Fatal("Invalid unit. Use 'C' for Celsius or 'F' for Fahrenheit.")
	}
}
