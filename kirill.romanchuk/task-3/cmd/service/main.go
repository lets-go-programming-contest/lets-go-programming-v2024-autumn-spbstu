package main

import (
	"github.com/kirill.romanchuk/task-3/strtructs/config"
	"github.com/kirill.romanchuk/task-3/strtructs/currencyRates"
)

func main() {
	/*
		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintln(os.Stderr, "Recovered from panic:", r)
				os.Exit(1)
			}
		}()
	*/
	cfg := config.Config{}
	cfg.Parse()

	currencyRates := currencyRates.CurrencyRates{}
	currencyRates.ParseXML(cfg.InputFile)

	currencyRates.SortByValue(true)

	currencyRates.ExportSelectedCurrencyRatesToJSON(cfg.OutputFile, "NumCode", "CharCode", "Value")
}
