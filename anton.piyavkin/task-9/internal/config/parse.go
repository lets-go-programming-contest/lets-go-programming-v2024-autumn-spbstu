package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func Load(name string) *Config {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	cfg := &Config{}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
