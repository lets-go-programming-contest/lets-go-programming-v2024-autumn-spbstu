package config

import (
	"os"
	structs "task-3/internal/structs"

	"gopkg.in/yaml.v3"
)

func ReadConfig() structs.Config {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var config structs.Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
