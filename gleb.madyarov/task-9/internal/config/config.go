package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Host       string `yaml:"host"`
	PortDb     string `yaml:"port_db"`
	DbUser     string `yaml:"dbUser"`
	Password   string `yaml:"password"`
	DbName     string `yaml:"dbName"`
	PortServer string `yaml:"port_server"`
}

func ReadConfig(path string) (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
