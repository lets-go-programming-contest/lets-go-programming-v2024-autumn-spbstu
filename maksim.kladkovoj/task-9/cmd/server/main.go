package main

import (
	"os"

	"github.com/Mmmakskl/task-9/cmd/server/internal/app"
	"github.com/Mmmakskl/task-9/cmd/server/internal/config"
)

func main() {
	configFile, err := os.Open(*ConfigPathFlag)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	cfg, err := config.Unmarshaller(configFile)
	if err != nil {
		panic(err)
	}

	app, err := app.NewApp(&cfg)
	if err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
