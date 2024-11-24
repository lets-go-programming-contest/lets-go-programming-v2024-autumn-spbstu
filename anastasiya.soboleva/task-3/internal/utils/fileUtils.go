package utils

import (
	"fmt"
	"os"
)

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	return file, nil
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(fmt.Errorf("error closing file: %w", err))
	}
}
