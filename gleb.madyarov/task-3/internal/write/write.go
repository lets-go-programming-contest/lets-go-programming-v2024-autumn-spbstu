package write

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Madyarov-Gleb/task-3/internal/read"
)

func WriteResult(data []byte, config read.Config) error {
	err := os.MkdirAll(filepath.Dir(config.Output), 0777)
	if err != nil {
		return fmt.Errorf("failed to create a directory: %w", err)
	}
	file, err := os.Create(config.Output)
	if err != nil {
		return fmt.Errorf("failed to create a file: %w", err)
	}
	defer file.Close()
	err = os.WriteFile(config.Output, data, 0666)
	if err != nil {
		return fmt.Errorf("failed to write to a file: %w", err)
	}
	return nil
}
