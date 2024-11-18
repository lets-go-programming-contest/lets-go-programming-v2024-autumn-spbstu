package data_reader

import (
	"fmt"
	"os"

	"github.com/solomonalfred/task-3/internal/schemas"
	"gopkg.in/yaml.v3"
)

func GetConfig() schemas.ConfigStruct {
	configPath := ConfigFlag()
	fmt.Println()
	configData, err := os.OpenFile(configPath, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer configData.Close()
	data := make([]byte, 1024)
	n, err := configData.Read(data)
	if err != nil {
		panic(err)
	}
	var config schemas.ConfigStruct
	err = yaml.Unmarshal(data[:n], &config)
	if err != nil {
		panic(err)
	}
	return config
}
