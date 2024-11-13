package valuteStrukts

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type CurrencyValue float64

func (c *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	s = strings.ReplaceAll(s, ",", ".")

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("failed to parse currency value: %w", err)
	}

	*c = CurrencyValue(value)
	return nil
}

type Valute struct {
	NumCode  string        `xml:"NumCode" json:"num_code"`
	CharCode string        `xml:"CharCode" json:"char_code"`
	Value    CurrencyValue `xml:"Value" json:"value"`
}

type ValuteRate struct {
	XMLName    xml.Name `xml:"ValCurs"`
	ValuteRate []Valute `xml:"Valute" json:"valute"`
}
