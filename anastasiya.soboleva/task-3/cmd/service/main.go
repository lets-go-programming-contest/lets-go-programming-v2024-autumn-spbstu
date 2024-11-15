package main

import (
	"flag"
	"fmt"
	"log"

	"anastasiya.soboleva/task-3/internal/parser"
	"anastasiya.soboleva/task-3/internal/sort"
	"anastasiya.soboleva/task-3/internal/utils"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error: %v\n", r)
		}
	}()
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()
	cfg, err := parser.ParseConfig(*configPath)
	if err != nil {
		panic(err)
	}
	rates, err := parser.ParseRates(cfg.InputFile)
	if err != nil {
		panic(err)
	}
	sort.RatesSort(rates)
	err = utils.SaveRates(rates, cfg.OutputFile)
	if err != nil {
		panic(err)
	}
	log.Println("Data processed and saved successfully.")
}
