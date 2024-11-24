package main

import (
	"github.com/EmptyInsid/task-3/configs/configData"
	"github.com/EmptyInsid/task-3/internal/app"
	"github.com/EmptyInsid/task-3/internal/parseFlag"
)

func main() {

	configPath, err := parseFlag.ParseFlags()
	if err != nil {
		panic(err)
	}

	config, err := configData.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	err = app.Run(config)
	if err != nil {
		panic(err)
	}
}
