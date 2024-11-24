package main

import (
	"github.com/hahapathetic/task-3/internal/config"
	"github.com/hahapathetic/task-3/internal/valuteProcessing"
	"github.com/hahapathetic/task-3/internal/valuteStructs"
)

func main() {
	data := valuteStructs.ValuteRate{}

	path := config.ReadFilePath()

	cfg := config.Config{}

	err := config.Parse(&cfg, path)
	if err != nil {
		panic(err)
	}

	err = valuteProcessing.ParseFromXML(&data, cfg.Input)
	if err != nil {
		panic(err)
	}

	valuteProcessing.Sort(data.ValuteRate)

	dataJSON, err := valuteProcessing.ParseToJSON(&data)
	if err != nil {
		panic(err)
	}

	err = valuteProcessing.WriteResult(dataJSON, cfg)
	if err != nil {
		panic(err)
	}
}
