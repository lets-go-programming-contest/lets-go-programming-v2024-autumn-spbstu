package main

import (
	"github.com/katagiriwhy/task-3/internal"
)

func main() {

	yaml := internal.ReadFlag()

	cfg := internal.Config{}
	curr := internal.Currencies{}

	err := internal.ReadAndParseConfig(&cfg, yaml)
	err = internal.ParseXml(&curr, cfg.InputFile)
	err = internal.Convert(&curr, cfg.OutputDir)

	panic(err)
}
