package main

import (
	config "task-3/internal/config"
	read "task-3/internal/readFile"
	write "task-3/internal/writeFile"
)

func main() {

	config, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}

	currencies, err := read.ParseXML(config.InputFile)
	if err != nil {
		panic(err)
	}

	read.SortValutes(&currencies)

	err = write.WriteToJson(currencies, config.OutputFile)
	if err != nil {
		panic(err)
	}
}
