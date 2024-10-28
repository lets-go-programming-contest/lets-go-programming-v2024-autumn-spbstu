package main

import (
	"fmt"
	"github.com/Piyavva/task-3/internal/config"
	"github.com/Piyavva/task-3/internal/input"
)

func main() {
	cfg := config.ParseConfig()
	valQuotes := input.ReadFile(cfg)
	fmt.Println(valQuotes)
}
