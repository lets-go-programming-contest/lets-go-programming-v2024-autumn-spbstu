package data_reader

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/solomonalfred/task-3/internal/schemas"
)

func GetConfig() (*schemas.ConfigStruct, error) {
	configPath := ConfigFlag()
	fmt.Println()
	configData, err := os.OpenFile(configPath, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
	}
	defer configData.Close()
	data := make([]byte, 1024)
	n, err := configData.Read(data)
	if err != nil {
		return nil, err
	}
	var config schemas.ConfigStruct
	err = yaml.Unmarshal(data[:n], &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
