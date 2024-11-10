package main

import (
	"task-3/internal/config"
	"task-3/internal/valuteProcessing"
	"task-3/internal/valuteStrukts"
)

func main() {
	data := valuteStrukts.ValuteRate{}
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
