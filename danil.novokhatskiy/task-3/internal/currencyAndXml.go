package internal

import (
	"encoding/xml"
	"os"
	"strings"
)

type Currency struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type Currencies struct {
	Currencies []Currency `xml:"Valute"`
}

func ParseXml(cur *Currencies, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	data = []byte(strings.ReplaceAll(string(data), ",", "."))
	err = xml.Unmarshal(data, &cur)
	if err != nil {
		return err
	}
	return nil
}
