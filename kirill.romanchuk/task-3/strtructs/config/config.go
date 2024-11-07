package config

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (c *Config) Parse() error {
	configFileParse := flag.String("config", "cmd/service/config.yaml", "Path to the configuration file")

	flag.Parse()

	data, err := os.ReadFile(*configFileParse)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}
