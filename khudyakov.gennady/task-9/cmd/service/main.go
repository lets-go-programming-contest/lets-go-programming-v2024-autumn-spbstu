package main

import (
	"flag"
	"fmt"

	"github.com/KRYST4L614/task-9/internal/app"
	"github.com/KRYST4L614/task-9/internal/config"
)

func main() {
	configPath := flag.String("config", "../../configs/config.yaml", "specify config path")
	flag.Parse()

	conf, err := config.ReadConfigFromYAML[app.Config](*configPath)
	if err != nil {
		panic(fmt.Errorf("read config from '%s' failed: %w", *configPath, err))
	}
	err = config.ValidateConfig(conf)
	if err != nil {
		panic(fmt.Errorf("'%v' parsing failed: %w", *configPath, err))
	}

	fmt.Println("Starting...")

	app, err := app.NewApp(conf)
	if err != nil {
		panic(err)
	}

	err = app.Start()
	if err != nil {
		panic(err)
	}
}
