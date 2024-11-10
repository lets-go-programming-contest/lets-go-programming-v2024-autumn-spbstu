package main

import (
	"flag"

	"erdem.istaev/task-3/internal/parser"
	"erdem.istaev/task-3/internal/sort_pkg"
	"erdem.istaev/task-3/internal/writer"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to config file")
	flag.Parse()

	config, err := parser.LoadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	valutes, err := parser.LoadValutes(config.InputFile)
	if err != nil {
		panic(err)
	}

	sort_pkg.SortValutes(valutes)

	if err = writer.SaveResults(valutes, config.OutputFile); err != nil {
		panic(err)
	}
}
