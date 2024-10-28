package config

import (
	"github.com/Piyavva/task-3/internal/structures"
	"gopkg.in/yaml.v3"
	"os"
)

func ParseConfig() structures.Config {
	data, err := os.ReadFile(NameFile)
	if err != nil {
		panic(err)
	}
	var config structures.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
