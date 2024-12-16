package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// перенести в tmp
type Config struct {
	DataBase DataBaseCfg
	Server   ServerCfg
}

type DataBaseCfg struct {
	UserDB     string `yaml:"dbUser"`
	PasswordDB string `yaml:"dbPassword"`
	NameDB     string `yaml:"dbName"`
	HostDB     string `yaml:"dbHost"`
	PortDB     int    `yaml:"dbPort"`
}

type ServerCfg struct {
	Port string `yaml:"port"`
}

func ParseConfig(path string) (*Config, error) {
	config := &Config{}

	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(yamlFile, config); err != nil {
		return nil, err
	}

	return config, nil
}
