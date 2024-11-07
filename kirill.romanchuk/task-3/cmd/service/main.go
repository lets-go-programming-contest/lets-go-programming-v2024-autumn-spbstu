package main

import (
	"fmt"
	"os"

	"github.com/kirill.romanchuk/task-3/strtructs/config"
	"github.com/kirill.romanchuk/task-3/strtructs/currencyRates"
)

func main() {
	cfg := config.Config{}

	err := cfg.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing config:", err)
		os.Exit(1)
	}

	currencyRates := currencyRates.CurrencyRates{}

	err = currencyRates.ParseXML(cfg.InputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing XML:", err)
		os.Exit(1)
	}

	currencyRates.SortByValue(true)

	err = currencyRates.ExportSelectedCurrencyRatesToJSON(cfg.OutputFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error exporting to JSON:", err)
		os.Exit(1)
	}
}
