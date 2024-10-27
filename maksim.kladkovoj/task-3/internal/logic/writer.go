package logic

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJSON(outPath string, valCurs *ValCurs) error {

	SortValutes(valCurs)

	jsonData, err := json.MarshalIndent(valCurs.Valutes, "", " ")
	if err != nil {
		return fmt.Errorf("Failure to marshall JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(outPath), 0777); err != nil {
		return fmt.Errorf("Failure make directory: %w", err)
	}

	file, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("Failure to create JSON file: %w", err)
	}
	defer file.Close()

	if err := os.WriteFile(outPath, jsonData, 0644); err != nil {
		return fmt.Errorf("Failure to write JSON file: %w", err)
	}

	return nil
}
