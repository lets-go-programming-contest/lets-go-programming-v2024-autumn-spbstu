package config

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	DBHost     string `yaml:"host"`
	DBPort     int    `yaml:"port"`
	DBUser     string `yaml:"user"`
	DBPassword string `yaml:"password"`
	DBName     string `yaml:"name"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type LoggerConfig struct {
	Mod string `yaml:"mod"`
}

type Config struct {
	DBCfg     DatabaseConfig `yaml:"database"`
	ServerCfg ServerConfig   `yaml:"server"`
	LoggerCfg LoggerConfig   `yaml:"logger"`
}

func LoadConfig(r io.Reader) (Config, error) {
	var cfg Config
	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("parse config error: %w", err)
	}

	return cfg, nil
}
