package currency

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
	"strings"
	"path/filepath"
	"golang.org/x/net/html/charset"
	"bytes"
)

type Currency struct {
	NumCode   int64   `xml:"NumCode" json:"num-code"`
	CharCode  string  `xml:"CharCode" json:"char-code"`
	Nominal   int64   `xml:"Nominal" json:"nominal"`
	Name      string  `xml:"Name" json:"name"`
	Value     float64 `xml:"Value" json:"value"`
	VunitRate float64 `xml:"VunitRate" json:"vunit-rate"`
}

type Currencies struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Name       string     `xml:"name,attr"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyJSON struct {
	NumCode   int64   `json:"num-code"`
	CharCode  string  `json:"char-code"`
	Nominal   int64   `json:"nominal"`
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	VunitRate float64 `json:"vunit-rate"`
}

func (currencies *Currencies) ParseXML(pathToXML string) {
	data, err := os.ReadFile(pathToXML)

	if err != nil {
		panic(err)
	}

	data = []byte(strings.ReplaceAll(string(data), ",", "."))

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&currencies)

	if err != nil {
		panic(err)
	}
}

func (currencies *Currencies) SortByValue(reverse bool) {
	if reverse {
		sort.Slice(currencies.Currencies, func(i, j int) bool {
			return currencies.Currencies[i].Value > currencies.Currencies[j].Value
		})
	} else {
		sort.Slice(currencies.Currencies, func(i, j int) bool {
			return currencies.Currencies[i].Value < currencies.Currencies[j].Value
		})
	}
}

func (currencies *Currencies) WriteCurrenciesToJSON(filename string, fields ...string) {
	var output []CurrencyJSON

	for _, currency := range currencies.Currencies {
		jsonCurrency := CurrencyJSON{
			NumCode:   currency.NumCode,
			CharCode:  currency.CharCode,
			Nominal:   currency.Nominal,
			Name:      currency.Name,
			Value:     currency.Value,
			VunitRate: currency.VunitRate,
		}

		if len(fields) > 0 {
			filteredCurrency := CurrencyJSON{}
			for _, field := range fields {
				switch field {
				case "NumCode":
					filteredCurrency.NumCode = jsonCurrency.NumCode
				case "CharCode":
					filteredCurrency.CharCode = jsonCurrency.CharCode
				case "Nominal":
					filteredCurrency.Nominal = jsonCurrency.Nominal
				case "Name":
					filteredCurrency.Name = jsonCurrency.Name
				case "Value":
					filteredCurrency.Value = jsonCurrency.Value
				case "VunitRate":
					filteredCurrency.VunitRate = jsonCurrency.VunitRate
				}
			}
			output = append(output, filteredCurrency)
		} else {
			output = append(output, jsonCurrency)
		}
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, jsonData, os.ModePerm)
	if err != nil {
		panic(err)
	}
}