package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"anastasiya.soboleva/task-3/internal/models"
)

func SaveRates(rates []models.Currency, outputPath string) {
	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		panic(err)
	}
	file, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	defer CloseFile(file)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(rates); err != nil {
		panic(err)
	}
}
