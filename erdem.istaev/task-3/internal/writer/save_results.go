package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"erdem.istaev/task-3/internal/parser"
)

func SaveResults(valutes []parser.Valute, outputFile string) error {
	if err := os.MkdirAll(filepath.Dir(outputFile), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	return encoder.Encode(valutes)
}
