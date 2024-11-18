package config

import (
	"os"
	errors "task-3/internal/errorsExt"
	structs "task-3/internal/readFile"

	"gopkg.in/yaml.v3"
)

func ReadConfig() (structs.Config, error) {
	var config structs.Config

	data, err := os.ReadFile(fileName)
	if err != nil {
		return config, errors.ErrorLocate(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, errors.ErrorLocate(err)
	}
	return config, nil
}
