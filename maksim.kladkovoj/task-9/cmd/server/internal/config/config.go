package config

import (
	"errors"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

var ErrParseConfig = errors.New("parse config failed")

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DbName   string `yaml:"dbname"`
	} `yaml:"database"`
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`
}

func Unmarshaller(r io.Reader) (Config, error) {
	var cfg Config

	if err := yaml.NewDecoder(r).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrParseConfig, err)
	}

	return cfg, nil
}
