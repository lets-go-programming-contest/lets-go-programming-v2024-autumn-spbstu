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

func (config *Config) ParseConfig() {
	configFileParse := flag.String("config", "../../config.yaml", "Path to the configuration file")
	flag.Parse()

	data, err := os.ReadFile(*configFileParse)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, config)

	if err != nil {
		panic(err)
	}
}
