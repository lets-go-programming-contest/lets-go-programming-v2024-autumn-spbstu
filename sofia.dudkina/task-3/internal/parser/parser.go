package parser

import (
	"bytes"
	"encoding/xml"
	"github.com/sssidkn/task-3/internal/entities"
	"golang.org/x/net/html/charset"
	"os"
)

func ParseFile(filePath string) *entities.ValCurs {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	valCurs := ParseXML(data)
	return valCurs
}

func ParseXML(data []byte) *entities.ValCurs {
	valCurs := new(entities.ValCurs)
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}
	return valCurs
}

func parseJSON() {

}
