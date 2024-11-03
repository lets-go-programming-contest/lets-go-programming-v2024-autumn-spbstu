package parser

import (
	"gopkg.in/yaml.v3"

	"anastasiya.soboleva/task-3/internal/models"
	"anastasiya.soboleva/task-3/internal/utils"
)

func ParseConfig(path string) *models.Configs {
	file := utils.OpenFile(path)
	defer utils.CloseFile(file)
	var cfg models.Configs
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}
