package main

import (
	processing "task-3/internal/ProcessingXML"
	"task-3/internal/config"
)

func main() {
	config, err := config.IsOkConfigFile()
	if err != nil {
		panic(err)
	}

	err = processing.ProcessingXML(config)
	if err != nil {
		panic(err)
	}

}
