package config

import (
	"flag"
)

func GetConfigPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "", "path to the configuration file")
	flag.Parse()
	return configPath
}
