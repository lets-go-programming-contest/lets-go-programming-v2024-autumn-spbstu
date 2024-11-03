package main

import (
	"flag"
	"log"

	"anastasiya.soboleva/task-3/internal/parser"
	"anastasiya.soboleva/task-3/internal/sort"
	"anastasiya.soboleva/task-3/internal/utils"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()
	cfg := parser.ParseConfig(*configPath)
	rates := parser.ParseRates(cfg.InputFile)
	sort.RatesSort(rates)
	utils.SaveRates(rates, cfg.OutputFile)
	log.Println("Данные обработаны и сохранены успешно.")
}
