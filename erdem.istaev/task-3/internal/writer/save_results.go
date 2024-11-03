package writer

import (
	"encoding/json"
	"os"

	"erdem.istaev/task-3/internal/parser"
)

func SaveResults(valutes []parser.Valute, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	return encoder.Encode(valutes)
}
