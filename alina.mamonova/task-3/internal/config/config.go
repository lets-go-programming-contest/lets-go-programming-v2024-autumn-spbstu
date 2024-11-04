package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func Parse(c *Config, f string) error {
	cContent, err := os.ReadFile(f)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = yaml.Unmarshal(cContent, c)
	if err != nil {
		return fmt.Errorf("failed to process file: %w", err)
	}

	return nil
}

func ReadFilePath() string {
	path := ""
	flag.StringVar(&path, "config", "", "the config file")
	flag.Parse()
	return string(path)
}
