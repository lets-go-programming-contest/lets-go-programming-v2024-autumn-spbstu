package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"task-3/internal/userErrors"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func IsOkConfigFile() (Config, error) {

	configFileName := flag.String("config", "config.yaml", "Config file")
	flag.Parse()
	if _, err := os.Stat(*configFileName); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("%w: %w", userErrors.ErrConfigFileIsNotExist, err)
	}

	dataOfConfig, err := os.ReadFile(*configFileName)
	if err != nil {
		return Config{}, fmt.Errorf("%w: %w", userErrors.ErrReadingFile, err)
	}

	var config Config
	err = yaml.Unmarshal(dataOfConfig, &config)
	if err != nil {
		return Config{}, fmt.Errorf("%w: %w", userErrors.ErrDeserializationFailure, err)
	}

	if _, err := os.Stat(config.InputFile); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("%w: %w", userErrors.ErrInputFileIsNotExist, err)
	}

	if _, err := os.Stat(config.OutputFile); os.IsNotExist(err) {
		outputDir := filepath.Dir(config.OutputFile)
		if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
			return Config{}, fmt.Errorf("%w: %w", userErrors.ErrMkdirFailure, err)
		}

		if config.OutputFile == "" {
			config.OutputFile = filepath.Join(outputDir, "output.json")
		}

		file, err := os.Create(config.OutputFile)
		if err != nil {
			return Config{}, fmt.Errorf("%w: %w", userErrors.ErrCreatingFileFailure, err)
		}
		defer file.Close()
	}

	return config, nil
}
