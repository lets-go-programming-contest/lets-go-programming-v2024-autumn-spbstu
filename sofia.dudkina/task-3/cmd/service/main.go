package main

import (
	"github.com/sssidkn/task-3/config"
	"github.com/sssidkn/task-3/internal/app"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
