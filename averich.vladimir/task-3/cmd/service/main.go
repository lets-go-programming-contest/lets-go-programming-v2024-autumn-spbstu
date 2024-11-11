package main

import (
    "task-3/internal/config"
    "task-3/internal/currency"
)

func main() {
	cfg := config.Config{}
	cfg.ParseConfig()

	currency := currency.Currencies{}
	currency.ParseXML(cfg.InputFile)

	currency.SortByValue(true)

	currency.WriteCurrenciesToJSON(cfg.OutputFile, "NumCode", "CharCode", "Value")
}