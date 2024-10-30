package main

import (
	"github.com/artem6554/task-3/config"
	"github.com/artem6554/task-3/read"
	"github.com/artem6554/task-3/write"
)

func main() {
	config := config.ReadConfig()
	currencies := read.ParseXML(config.InputFile)
	read.SortValutes(&currencies)
	write.WriteToJson(currencies, config.OutputFile)

}
