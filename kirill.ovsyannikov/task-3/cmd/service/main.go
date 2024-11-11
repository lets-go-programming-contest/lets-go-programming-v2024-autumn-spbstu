package main

import (
	config "task-3/internal/config"
	read "task-3/internal/readFile"
	write "task-3/internal/writeFile"
)

func main() {
	config := config.ReadConfig()
	currencies := read.ParseXML(config.InputFile)
	read.SortValutes(&currencies)
	write.WriteToJson(currencies, config.OutputFile)

}
