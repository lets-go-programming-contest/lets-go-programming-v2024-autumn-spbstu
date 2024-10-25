package internal

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadFlag() string {
	loc := ""
	flag.StringVar(&loc, "cfg", "", "Config to read")
	flag.Parse()
	return loc
}

type Config struct {
	InputFile string `yaml:"input-file"`
	OutputDir string `yaml:"output-file"`
}

func ReadAndParseConfig(cfg *Config, f string) error {
	file, err := os.ReadFile(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(file, cfg)
	if err != nil {
		return err
	}
	return nil
}
