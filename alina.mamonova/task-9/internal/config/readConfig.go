package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type DbData struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type ServerData struct {
	Addr string `yaml:"adress"`
}

type ConfigPaths struct {
	DBConfig     string `yaml:"dbconfig"`
	ServerConfig string `yaml:"serverconfig"`
}

func ReadDbConfig() DbData {
	data, err := os.ReadFile(dbFileName)
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
	data, err := os.ReadFile(serverFileName)
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

func loadConfigPaths(filePath string) (*ConfigPaths, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var configPaths ConfigPaths
	if err := yaml.Unmarshal(data, &configPaths); err != nil {
		return nil, err
	}

	return &configPaths, nil
}
