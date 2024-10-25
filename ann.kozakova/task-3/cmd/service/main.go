package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	configPath string
)

func initFlag() {
	flag.StringVar(&configPath, "config", "config.yml", "The path to the configuration file")
	flag.Parse()
}

type Config struct {
	Input  string `yaml:"input-file"`
	Output string `yaml:"output-file"`
}

func main() {
	initFlag()
	configFile, err := os.OpenFile(configPath, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	buffer := make([]byte, 512)
	n, err := configFile.Read(buffer)
	if err != nil {
		panic(err)
	}

	var config Config
	err = yaml.Unmarshal(buffer[:n], &config)
	if err != nil {
		panic(err)
	}

	fmt.Println("input: ", config.Input)
	fmt.Println("output: ", config.Output)

	inputFile, err := os.OpenFile(config.Input, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
}
