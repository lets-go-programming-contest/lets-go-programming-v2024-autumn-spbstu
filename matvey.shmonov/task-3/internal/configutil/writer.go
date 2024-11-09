package configutil

import (
	"fmt"
	"os"
	"path/filepath"
)

func WriteResult(data []byte, config Config) error {
	err := os.MkdirAll(filepath.Dir(config.OutFilename), 0777)
	if err != nil {
		return fmt.Errorf("failed to create a directory: %w", err)
	}
	file, err := os.Create(config.OutFilename)
	if err != nil {
		return fmt.Errorf("failed to create a file: %w", err)
	}
	defer file.Close()
	err = os.WriteFile(config.OutFilename, data, 0666)
	if err != nil {
		return fmt.Errorf("failed to write to a file: %w", err)
	}
	return nil
}
