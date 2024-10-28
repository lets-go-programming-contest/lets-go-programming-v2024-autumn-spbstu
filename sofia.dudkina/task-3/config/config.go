package config

import (
	"errors"
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	var filePath string
	flag.StringVar(&filePath, "config", "root", "path to configuration file")
	flag.Parse()

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return nil, errors.New("the configuration file does not match the project configuration structure")
	}
	if _, err := os.Stat(cfg.InputFile); os.IsNotExist(err) {
		return nil, err
	}
	return cfg, nil
}
