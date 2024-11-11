package src

import (
	"flag"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/solomonalfred/task-3/internal/structures"
)

var configPath string

func initFlag() string {
	flag.StringVar(&configPath, "config", "config.yml", "Configuration path")
	flag.Parse()
	return configPath
}

func ParseConfig() structures.ConfigStruct {
	configPath := initFlag()
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

	var config structures.ConfigStruct
	err = yaml.Unmarshal(buffer[:n], &config)
	if err != nil {
		panic(err)
	}
	return config
}


