package main

import (
	"log"

	"github.com/IDevFrye/task-3/internal/config"
	"github.com/IDevFrye/task-3/internal/currency"
	errProc "github.com/IDevFrye/task-3/internal/errors"
)

func main() {
	configPath := config.GetConfigPath()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(errProc.ErrConfigNotFound)
	}

	if err := currency.DownloadCurrencyData(cfg.InputFile); err != nil {
		panic(errProc.ErrDataDownload)
	}

	data, err := currency.FetchCurrencyData(cfg.InputFile)
	if err != nil {
		panic(errProc.ErrInvalidXMLFormat)
	}

	processedData := currency.ProcessCurrencies(data)

	if err := currency.SaveAsJSON(processedData, cfg.OutputFile); err != nil {
		panic(errProc.ErrDataWriting)
	}

	log.Println("Currency data successfully processed and saved.")
}
