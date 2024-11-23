package parser

import (
	"anastasiya.soboleva/task-3/internal/models"
	"anastasiya.soboleva/task-3/internal/utils"
)

func ParseRates(filePath string) ([]models.Currency, error) {
	file, err := utils.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer utils.CloseFile(file)
	valCurs, err := parseXML(file)
	if err != nil {
		return nil, err
	}
	currencies, err := ConvertToModels(valCurs)
	if err != nil {
		return nil, err
	}
	return currencies, nil
}
