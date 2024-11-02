package parser

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/sssidkn/task-3/internal/entities"
	"golang.org/x/net/html/charset"
)

func ParseFile(path string) (*entities.ValCurs, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	extension := filepath.Ext(path)
	pattern := `([\d.,]+)`
	re := regexp.MustCompile(pattern)
	replaceFunc := func(str string) string {
		return strings.Replace(str, ",", ".", 1)
	}
	data = []byte(re.ReplaceAllStringFunc(string(data), replaceFunc))
	cursData := &entities.ValCurs{}
	switch extension {
	case ".xml":
		cursData, err = parseXML(data)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported file extension: %s", extension)
	}
	return cursData, err
}

func parseXML(data []byte) (*entities.ValCurs, error) {
	cursData := new(entities.ValCurs)
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&cursData)
	if err != nil {
		return nil, err
	}
	return cursData, nil
}
