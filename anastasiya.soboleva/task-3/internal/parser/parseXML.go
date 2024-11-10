package parser

import (
	"encoding/xml"
	"golang.org/x/net/html/charset"
	"io"

	"anastasiya.soboleva/task-3/internal/models"
)

func parseXML(reader io.Reader) (models.ValCurs, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	var valCurs models.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, err
	}
	return valCurs, nil
}
