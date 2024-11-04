package configData

import (
	"os"

	"github.com/EmptyInsid/task-3/internal/errorUtils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}

	return config, nil
}
