package internal

import (
	"encoding/xml"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

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
	content := strings.ReplaceAll(string(data), `encoding="windows-1251"`, `encoding="UTF-8"`)
	content = strings.ReplaceAll(content, ",", ".")
	decoder := charmap.Windows1251.NewDecoder()
	utf8Content, _, err := transform.String(decoder, content)
	if err != nil {
		return err
	}
	xmlDecoder := xml.NewDecoder(strings.NewReader(utf8Content))
	err = xmlDecoder.Decode(&cur)
	if err != nil {
		return err
	}
	return nil
}
