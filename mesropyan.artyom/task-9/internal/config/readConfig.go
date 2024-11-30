package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type DbData struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func ReadDbConfig() DbData {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var config DbData
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
