package writeJson

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/EmptyInsid/task-3/internal/errorUtils"
	"github.com/EmptyInsid/task-3/internal/structs/xmlStruct"
)

func ProcessJson(xmlData *xmlStruct.DataCursStruct, outputFilePath string) error {
	file, err := OpenFilePath(outputFilePath)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}
	defer file.Close()

	err = WriteJson(xmlData, file)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}

	return nil
}

func OpenFilePath(outputFilePath string) (*os.File, error) {
	ext := filepath.Ext(outputFilePath)
	if ext != ".json" {
		return nil, errorUtils.ErrorWithLocation(fmt.Errorf("wrong file extantion: %s", ext))
	}

	_, err := os.Stat(outputFilePath)
	if err != nil {
		if err = createDir(outputFilePath); err != nil {
			return nil, errorUtils.ErrorWithLocation(err)
		}
	}

	file, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}
	return file, nil
}

func WriteJson(xmlData *xmlStruct.DataCursStruct, file *os.File) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(xmlData); err != nil {
		return errorUtils.ErrorWithLocation(err)
	}
	return nil
}

func createDir(outputFilePath string) error {
	err := os.MkdirAll(filepath.Dir(outputFilePath), 0644)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}
	return nil
}
