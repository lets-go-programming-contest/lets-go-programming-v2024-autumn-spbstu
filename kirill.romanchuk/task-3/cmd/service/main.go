package main

import (
	"fmt"

	"github.com/kirill.romanchuk/task-3/strtructs/config"
	"github.com/kirill.romanchuk/task-3/strtructs/currencyRates"
)

func main() {
	config := config.Config{}

	err := config.Parse()

	if err != nil {
		fmt.Print(err)
	}
	currencyRates := currencyRates.CurrencyRates{}
	currencyRates.ParseXML(config.InputFile)
	currencyRates.SortByValue(true)
	currencyRates.ExportSelectedCurrencyRatesToJSON(config.OutputFile)
}
