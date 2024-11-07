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
	}
	currencyRates := strtructs.CurrencyRates{}
	currencyRates.ParseXML(config.InputFile)
	currencyRates.SortByValue(true)
	currencyRates.ExportSelectedCurrencyRatesToJSON(config.OutputFile)
}
