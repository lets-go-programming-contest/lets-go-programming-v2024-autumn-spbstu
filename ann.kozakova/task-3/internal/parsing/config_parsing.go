package parsing

import (
	"github.com/nutochk/task-3/internal/flag"
	"github.com/nutochk/task-3/internal/structures"
	"gopkg.in/yaml.v3"
	"os"
)

func ConfigParsing() structures.Config {
	configPath := flag.InitFlag()
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

	var config structures.Config
	err = yaml.Unmarshal(buffer[:n], &config)
	if err != nil {
		panic(err)
	}
	return config
}
