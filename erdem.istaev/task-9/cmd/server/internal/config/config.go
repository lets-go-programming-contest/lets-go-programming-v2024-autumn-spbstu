package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"

	"erdem.istaev/task-9/internal/repository/database"
)

var (
	ErrParseConfig = errors.New("error parsing config")
)

type Config struct {
	ServerPort string            `yaml:"host"`
	DB         database.DBConfig `yaml:"db"`
}

func Unmarshall(r io.Reader) (Config, error) {
	var cfg Config

	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrParseConfig, err)
	}

	return cfg, nil
}
