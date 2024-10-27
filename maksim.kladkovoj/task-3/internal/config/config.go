package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func LoadConfig(path string, conf *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error open files: %w", err)
	}

	if err := yaml.Unmarshal(data, conf); err != nil {
		return fmt.Errorf("Failed decode: %w", err)
	}

	return nil
}
