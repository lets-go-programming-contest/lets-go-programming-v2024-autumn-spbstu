package config

import (
	"flag"
)

func GetConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "configs/config.yml", "path to the configuration file")
	flag.Parse()
	return configPath
}
