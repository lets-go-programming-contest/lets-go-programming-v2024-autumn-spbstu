package config

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func NewConfig() *Config {
	cfg := &Config{}
	var filePath string
	flag.StringVar(&filePath, "config", "root", "path to configuration file")
	flag.Parse()

	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	if cfg.InputFile == "" || cfg.OutputFile == "" {
		panic(errors.New("the configuration file does not match the project configuration structure"))
	}
	if _, err := os.Stat(cfg.InputFile); os.IsNotExist(err) {
		panic(err)
	}
	return cfg
}
