package utils

import (
	"fmt"
	"os"
)

func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Printf("%v\n", err)
	}
}
