package main

import (
	"github.com/artem6554/task-3/config"
	"github.com/artem6554/task-3/read"
)

func main() {
	config := config.ReadConfig()
	read.ParseXML(config.InputFile)
}
