package read

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func ReadConfig() Config {
	pathConfig := flag.String("config", "", "path to config file")
	flag.Parse()

	var config Config
	yamlFile, err := os.ReadFile(*pathConfig)
	if err != nil {
		panic("the file path is missing")
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic("it was not possible to convert YAML into a structure")
	}
	return config
}
