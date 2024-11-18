package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"gopkg.in/yaml.v3"

	"task-3/internal/config"
	"task-3/internal/currency"
)

func main() {
	cfg := config.Config{}

	configFileParse := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	data, err := os.ReadFile(*configFileParse)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %s", *configFileParse)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Не удалось закодировать файл: %s", *configFileParse)
	}

	currencies, err := currency.ParseXML(cfg.InputFile)
	if err != nil {
		log.Fatalf("Не удалось распарсить XML-файл: %s", cfg.InputFile)
	}

	sort.Slice(currencies.Currencies, func(i, j int) bool {
		return currencies.Currencies[i].Value < currencies.Currencies[j].Value
	})

	err = currency.WriteCurrenciesToJSON(cfg.OutputFile, []string{"NumCode", "CharCode", "Value"}, cfg.InputFile)
	if err != nil {
		log.Fatalf("Не удалось записать данные в JSON-файл: %s", cfg.OutputFile)
	}

	fmt.Println("Данные успешно записаны в JSON-файл:", cfg.OutputFile)
}