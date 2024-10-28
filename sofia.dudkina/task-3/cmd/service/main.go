package main

import (
	"github.com/sssidkn/task-3/config"
	"github.com/sssidkn/task-3/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
