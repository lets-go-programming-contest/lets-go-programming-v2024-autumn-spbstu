package internal

import (
	"bytes"
	"encoding/xml"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type CurrencyRaw struct {
	NumCode  string `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    string `xml:"Value" json:"value"`
}

type Currency struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type Currencies struct {
	Currencies []Currency `xml:"Valute"`
}

type CurrenciesRaw struct {
	Currencies []CurrencyRaw `xml:"Valute"`
}

func ParseXml(cur *Currencies, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	var tmp CurrenciesRaw
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&tmp)
	cur.Currencies = make([]Currency, len(tmp.Currencies))
	for i := 0; i < len(tmp.Currencies); i++ {
		value, err := strconv.ParseFloat(strings.ReplaceAll(tmp.Currencies[i].Value, ",", "."), 64)
		if err != nil {
			return err
		}
		cur.Currencies[i] = Currency{
			NumCode:  tmp.Currencies[i].NumCode,
			CharCode: tmp.Currencies[i].CharCode,
			Value:    value,
		}
	}
	if err != nil {
		return err
	}
	return nil
}
