package processingXML

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
	"task-3/internal/config"
	"task-3/internal/userErrors"

	"golang.org/x/text/encoding/charmap"
)

type Valute struct {
	NumCode  string `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    string `xml:"Value" json:"value"`
}

type ValCurs struct {
	Valutes []Valute `xml:"Valute" json:"valutes"`
}

func ProcessingXML(config config.Config) error {
	dataWindows1251, err := os.ReadFile(config.InputFile)
	if err != nil {
		return fmt.Errorf("%w: %w", userErrors.ErrReadingFile, err)
	}

	decoder := charmap.Windows1251.NewDecoder()
	utf8Reader := decoder.Reader(strings.NewReader(string(dataWindows1251)))

	xmlDecoder := xml.NewDecoder(utf8Reader)
	xmlDecoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var valCurs ValCurs
	if err := xmlDecoder.Decode(&valCurs); err != nil {
		return fmt.Errorf("%w: %w", userErrors.ErrDeserializationFailure, err)
	}

	for i := range valCurs.Valutes {
		valCurs.Valutes[i].Value = strings.Replace(valCurs.Valutes[i].Value, ",", ".", 1)
	}

	sort.Slice(valCurs.Valutes, func(i, j int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	jsonData, err := json.MarshalIndent(valCurs, "", "    ")
	if err != nil {
		return fmt.Errorf("%w: %w", userErrors.ErrSerializationFailure, err)
	}

	if err := os.WriteFile(config.OutputFile, jsonData, os.ModePerm); err != nil {
		return fmt.Errorf("%w: %w", userErrors.ErrWriteFile, err)
	}

	return nil
}
