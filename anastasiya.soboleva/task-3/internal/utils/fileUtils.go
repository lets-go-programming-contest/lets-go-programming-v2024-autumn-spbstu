package utils

import (
	"os"
)

func OpenFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	return file
}

func CloseFile(file *os.File) {
	if err := file.Close(); err != nil {
		panic(err)
	}
}
