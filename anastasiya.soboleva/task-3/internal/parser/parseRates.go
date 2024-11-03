package parser

import (
	"anastasiya.soboleva/task-3/internal/models"
	"anastasiya.soboleva/task-3/internal/utils"
)

func ParseRates(filePath string) []models.Currency {
	file := utils.OpenFile(filePath)
	defer utils.CloseFile(file)
	valCurs := parseXML(file)
	currencies := ConvertToModels(valCurs)
	return currencies
}
