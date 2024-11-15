package config

import (
	"errors"
	"flag"
)

func ParseFlag() (string, error) {
	configPath := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	if *configPath == "" {
		return "", errors.New("Config file not found")
	}

	return *configPath, nil
}
