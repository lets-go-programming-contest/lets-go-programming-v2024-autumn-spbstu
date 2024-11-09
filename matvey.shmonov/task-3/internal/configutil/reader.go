package configutil

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InFilename  string `yaml:"input-file"`
	OutFilename string `yaml:"output-file"`
}

func ReadConfig() (Config, error) {
	pathConfig := flag.String("config", "", "path to config file")
	flag.Parse()

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
