package main

import (
	"github.com/Mmmakskl/task-3/internal/config"
	"github.com/Mmmakskl/task-3/internal/run"
)

func main() {

	configPath, err := config.ParseFlag()
	if err != nil {
		panic(err)
	}

	if err := run.Run(configPath); err != nil {
		panic(err)
	}
}
