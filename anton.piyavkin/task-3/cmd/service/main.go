package main

import (
	"fmt"
	"github.com/Piyavva/task-3/internal/input"
	"github.com/Piyavva/task-3/internal/parse"
)

func main() {
	cfg := parse.ParseConfig()
	valQuotes := input.ReadFile(cfg)
	input.SortCurrencies(&valQuotes)
	fmt.Println(valQuotes)
}
