package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadConfig() Config {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(data, config)
	if err != nil {
		panic(err)
	}
	return config
}
