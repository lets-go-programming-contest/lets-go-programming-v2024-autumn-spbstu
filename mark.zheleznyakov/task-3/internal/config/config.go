package config

import (
	"gopkg.in/yaml.v3"
)

type ConfigFile struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ParseFromString(s string, c *ConfigFile) error {
	return yaml.Unmarshal([]byte(s), c)
}
