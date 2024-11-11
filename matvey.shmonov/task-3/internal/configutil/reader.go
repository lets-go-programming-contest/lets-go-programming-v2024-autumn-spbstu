package configutil

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InFilename  string `yaml:"input-file"`
	OutFilename string `yaml:"output-file"`
}

func ReadConfig(pathConfig *string) (Config, error) {
	var config Config
	yamlFile, err := os.ReadFile(*pathConfig)
	if err != nil {
		return config, fmt.Errorf("the file path is missing: %w", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return config, fmt.Errorf("it was not possible to convert YAML into a structure: %w", err)
	}
	return config, nil
}
