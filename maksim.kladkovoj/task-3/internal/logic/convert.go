package logic

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

func Parser(filePath string, conf *ValCurs) error {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failure open file: %w", err)
	}

	file = []byte(strings.ReplaceAll(string(file), ",", "."))

	if err = xml.Unmarshal(file, conf); err != nil {
		return fmt.Errorf("Failure decoding file: %w", err)
	}

	return nil
}
