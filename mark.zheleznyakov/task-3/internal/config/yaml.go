package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Parse(c *ConfigFile, f string) error {
	cContent, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("your config file is cooked: %w", err)
	}

	err = yaml.Unmarshal(cContent, c)
	if err != nil {
		return fmt.Errorf("your config file contents are cooked: %w", err)
	}

	return nil
}
