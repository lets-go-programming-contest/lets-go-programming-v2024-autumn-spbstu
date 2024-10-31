package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/sssidkn/task-3/internal/entities"
	"golang.org/x/net/html/charset"
)

func ParseFile(path string) (*entities.OutputData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	extension := filepath.Ext(path)
	outData := &entities.OutputData{}
	switch extension {
	case ".xml":
		cursData := &entities.CursData{}
		cursData, err = parseXML(data)
		if err != nil {
			return nil, err
		}
		outData, err = dataTransfer(cursData)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", extension)
	}
	return outData, err
}

func parseXML(data []byte) (*entities.CursData, error) {
	cursData := new(entities.CursData)
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&cursData)
	if err != nil {
		return nil, err
	}
	return cursData, nil
}

func dataTransfer(cursData *entities.CursData) (*entities.OutputData, error) {
	outData := &entities.OutputData{}
	for i := range cursData.Valute {
		replacedValue := strings.ReplaceAll(cursData.Valute[i].Value, ",", ".")
		value, err := strconv.ParseFloat(replacedValue, 64)
		if err != nil {
			return nil, err
		}
		outData.Valute = append(outData.Valute, struct {
			NumCode  int     `json:"num_code,omitempty"`
			CharCode string  `json:"char_code,omitempty"`
			Value    float64 `json:"value,omitempty"`
		}{NumCode: cursData.Valute[i].NumCode, CharCode: cursData.Valute[i].CharCode, Value: value})
	}
	return outData, nil
}
