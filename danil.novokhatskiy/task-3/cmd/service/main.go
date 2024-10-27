package main

import (
	"log"

	"github.com/katagiriwhy/task-3/internal"
)

func main() {

	yaml := internal.ReadFlag()

	cfg := internal.Config{}
	curr := internal.Currencies{}

	err := internal.ReadAndParseConfig(&cfg, yaml)
	if err != nil {
		log.Fatal(err)
	}
	err = internal.ParseXml(&curr, cfg.InputFile)
	if err != nil {
		log.Fatal(err)
	}
	err = internal.Convert(&curr, cfg.OutputDir)
	if err != nil {
		log.Fatal(err)
	}
}
