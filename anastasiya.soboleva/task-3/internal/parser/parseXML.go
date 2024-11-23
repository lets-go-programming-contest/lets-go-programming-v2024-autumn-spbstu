package parser

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"io"

	"anastasiya.soboleva/task-3/internal/models"
)

func parseXML(reader io.Reader) (models.ValCurs, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	var valCurs models.ValCurs
	if err := decoder.Decode(&valCurs); err != nil {
		return valCurs, fmt.Errorf("error decoding XML: %w", err)
	}
	return valCurs, nil
}
