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

type ServerData struct {
	Addr string `yaml:"adress"`
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

func ReadServerConfig() ServerData {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	var config ServerData
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	return config
}
