package main

import (
	"github.com/Piyavva/task-3/internal/input"
	"github.com/Piyavva/task-3/internal/output"
	"github.com/Piyavva/task-3/internal/parse"
)

func main() {
	cfg := parse.ParseConfig()
	valQuotes := input.ReadFile(cfg)
	input.SortCurrencies(&valQuotes)
	output.Write(valQuotes, cfg)
}
