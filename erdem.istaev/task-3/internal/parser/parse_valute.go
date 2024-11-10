package parser

import (
	"io"
	"os"
	"strings"

	"encoding/xml"
	"golang.org/x/net/html/charset"
)

type Valute struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func LoadValutes(inputFile string) ([]Valute, error) {
	file, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	data = []byte(strings.ReplaceAll(string(data), ",", "."))

	decoder := xml.NewDecoder(strings.NewReader(string(data)))
	decoder.CharsetReader = charset.NewReaderLabel
	var valCurs ValCurs
	if err = decoder.Decode(&valCurs); err != nil {
		return nil, err
	}

	return valCurs.Valutes, nil
}
