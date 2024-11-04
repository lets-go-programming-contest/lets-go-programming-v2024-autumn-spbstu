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

	if err := currency.DownloadCurrencyData(cfg.InputFile); err != nil {
		log.Fatalf("Failed to download currency data: %v", err)
	}

	data, err := currency.FetchCurrencyData(cfg.InputFile)
	if err != nil {
		log.Fatalf("Failed to fetch currency data: %v", err)
	}

	processedData := currency.ProcessCurrencies(data)

	if err := currency.SaveAsJSON(processedData, cfg.OutputFile); err != nil {
		log.Fatalf("Failed to save processed data: %v", err)
	}

	log.Println("Currency data successfully processed and saved.")
}
