package main

import (
	"os"

	"erdem.istaev/task-9/cmd/server/internal/app"
	"erdem.istaev/task-9/cmd/server/internal/config"
	"erdem.istaev/task-9/internal/repository/database"
)

func main() {
	configFile, err := os.Open(*ConfigPathFlag)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	cfg, err := config.Unmarshall(configFile)
	if err != nil {
		panic(err)
	}

	newApp, db := app.NewApp(cfg)
	newApp.Run()
	defer newApp.Close()

	err = database.CloseDB(db)
	if err != nil {
		panic(err)
	}
}
