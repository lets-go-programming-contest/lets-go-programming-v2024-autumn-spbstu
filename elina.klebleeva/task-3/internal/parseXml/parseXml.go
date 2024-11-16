package parseXml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/EmptyInsid/task-3/internal/errorUtils"
	"github.com/EmptyInsid/task-3/internal/structs/xmlStruct"
	"golang.org/x/net/html/charset"
)

func ProcessXml(xmlfile string) (*xmlStruct.DataCursStruct, error) {

	xmlData, err := ReadXml(xmlfile)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}

	xmlStruct, err := ParseXml(xmlData)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}

	return xmlStruct, nil
}

func ReadXml(xmlfile string) ([]byte, error) {

	ext := filepath.Ext(xmlfile)
	if ext != ".xml" {
		return nil, errorUtils.ErrorWithLocation(fmt.Errorf("wrong file extantion: %s", ext))
	}

	xmlFile, err := os.Open(xmlfile)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}
	defer xmlFile.Close()

	xmlByteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}
	return xmlByteValue, nil

}

func ParseXml(xmlData []byte) (*xmlStruct.DataCursStruct, error) {

	xmlStruct := new(xmlStruct.DataCursStruct)

	decoder := xml.NewDecoder(bytes.NewReader(xmlData))
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&xmlStruct.ValCurs)
	if err != nil {
		return nil, errorUtils.ErrorWithLocation(err)
	}
	return xmlStruct, nil
}
