package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/EmptyInsid/task-9/cmd/app"
	"github.com/EmptyInsid/task-9/internal/config"
	myLog "github.com/EmptyInsid/task-9/internal/log"
)

func main() {
	CfigPathFlag := flag.String("config", "../../configs/config.yml", "Path to YAML config")
	flag.Parse()

	// Open config file
	configFile, err := os.Open(*CfigPathFlag)
	if err != nil {
		panic(err)
	}

	// Load date from config file
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}

	// Setup logger
	logger := myLog.Setup(cfg.LoggerCfg)
	logger.Info("initializing server", slog.String("address", cfg.ServerCfg.Port))
	logger.Debug("logger debug mode enabled")

	// Make new app with config
	app, err := app.NewApp(logger, cfg)
	if err != nil {
		logger.Error("create app", "error", err)
		panic(err)
	}

	logger.Info("success create new app")

	// Start app
	if err := app.Run(); err != nil {
		logger.Error("run app", "error", err)
		panic(err)
	}
}
