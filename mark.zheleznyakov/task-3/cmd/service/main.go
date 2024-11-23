package main

import (
	"log"

	"github.com/mrqiz/task-3/internal/config"
	"github.com/mrqiz/task-3/internal/currencies"
)

func main() {
	cLocation := config.ReadConfigFlag()

	cfg := config.ConfigFile{}
	crcs := currencies.Currencies{}

	err := config.Parse(&cfg, cLocation)
	if err != nil {
		log.Panicf("err: %v", err)
	}

	err = currencies.Parse(&crcs, cfg.InputFile)
	if err != nil {
		log.Panicf("err: %v", err)
	}

	err = currencies.Convert(&crcs, cfg.OutputFile)
	if err != nil {
		log.Panicf("err: %v", err)
	}
}
