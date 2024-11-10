package config

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (config *Config) GetConfig(data []byte) (Config, error) {
	err := yaml.Unmarshal(data, config)
	if err != nil {
		return *config, err
	}
	return *config, nil
}
