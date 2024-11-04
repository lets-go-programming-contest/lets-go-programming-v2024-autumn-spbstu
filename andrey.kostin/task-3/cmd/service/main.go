package main

import (
	"log"

	"github.com/IDevFrye/task-3/internal/config"
	"github.com/IDevFrye/task-3/internal/currency"
)

func main() {
	configPath := config.GetConfigPath()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	data, err := currency.FetchCurrencyData(cfg.InputFile)
	if err != nil {
		log.Fatalf("Failed to fetch currency data: %v", err)
	}

	processedData := currency.ProcessCurrencies(data)

	log.Println("Currency data successfully processed and saved.")
}
