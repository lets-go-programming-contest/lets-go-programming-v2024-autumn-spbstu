package main

import (
	"fmt"

	"github.com/kirill.romanchuk/task-3/strtructs"
)

func main() {
	config := strtructs.Config{}

	err := config.Parse()

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(config.InputFile)
		fmt.Println(config.OutputFile)
	}

	currencyRates := strtructs.CurrencyRates{}
	currencyRates.ParseXML(config.InputFile)
	fmt.Println(currencyRates.Date)
	for _, currency := range currencyRates.Currencies {
		fmt.Printf("NumCode: %d, CharCode: %s, Nominal: %d, Name: %s, Value: %.4f\n",
			currency.NumCode, currency.CharCode, currency.Nominal, currency.Name, currency.Value)
	}
	fmt.Println("////////////////////////////////////////////////////////////////////")
	currencyRates.SortByValue(true)
	for _, currency := range currencyRates.Currencies {
		fmt.Printf("NumCode: %d, CharCode: %s, Nominal: %d, Name: %s, Value: %.4f\n",
			currency.NumCode, currency.CharCode, currency.Nominal, currency.Name, currency.Value)
	}
}
