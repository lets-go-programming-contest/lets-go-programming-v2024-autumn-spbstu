package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"erdem.istaev/task-3/internal/parser"
	"erdem.istaev/task-3/internal/sort_pkg"
	"erdem.istaev/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	config, err := parser.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.MkdirAll(filepath.Dir(config.OutputFile), os.ModePerm); err != nil {
		log.Fatalf("Error creating output directory: %v", err)
	}

	valutes, err := parser.LoadValutes(config.InputFile)
	if err != nil {
		log.Fatalf("Error loading valutes: %v", err)
	}

	sort_pkg.SortValutes(valutes)

	if err := writer.SaveResults(valutes, config.OutputFile); err != nil {
		log.Fatalf("Error saving results: %v", err)
	}
}
