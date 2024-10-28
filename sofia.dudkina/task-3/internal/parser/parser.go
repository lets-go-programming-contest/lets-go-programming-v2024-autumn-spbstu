package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sssidkn/task-3/internal/entities"
	"golang.org/x/net/html/charset"
)

func ParseFile(path string) (*entities.CursData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data = bytes.ReplaceAll(data, []byte(","), []byte("."))
	extension := filepath.Ext(path)
	cursData := &entities.CursData{}
	switch extension {
	case ".xml":
		cursData, err = ParseXML(data)
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", extension)
	}
	return cursData, err
}

func ParseXML(data []byte) (*entities.CursData, error) {
	cursData := new(entities.CursData)
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&cursData.ValCurs)
	if err != nil {
		return nil, err
	}
	return cursData, nil
}
