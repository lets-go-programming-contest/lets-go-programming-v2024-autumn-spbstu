package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	Database DBConfig     `yaml:"database"`
	Server   ServerConfig `yaml:"server"`
}

func ReadConfig() (Config, error) {
	var config Config

	data, err := os.ReadFile(fileName)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
