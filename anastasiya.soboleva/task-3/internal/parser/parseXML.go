package parser

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"os"

	"anastasiya.soboleva/task-3/internal/models"
)

func parseXML(file *os.File) models.ValCurs {
	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel
	var valCurs models.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		panic(err)
	}
	return valCurs
}
