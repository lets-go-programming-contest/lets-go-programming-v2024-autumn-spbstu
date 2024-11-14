package currencyRates

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type CurrencyRates struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Name       string     `xml:"name,attr"` // Name of the currency market from XML.
	Currencies []Currency `xml:"Valute"`
}

func (c *CurrencyRates) ParseXML(pathToXML string) {
	file, err := os.ReadFile(pathToXML)
	if err != nil {
		panic(err)
	}

	if filepath.Ext(pathToXML) != ".xml" {
		panic(fmt.Errorf("файл по пути '%s' не является .xml", pathToXML))
	}

	modifiedContent := []byte(strings.ReplaceAll(string(file), ",", "."))

	dec := xml.NewDecoder(bytes.NewReader(modifiedContent))
	dec.CharsetReader = func(encoding string, input io.Reader) (io.Reader, error) {
		switch encoding {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("неподдерживаемая кодировка: %s", encoding)
		}
	}

	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
}

// If reverse is false, the currencies are sorted in ascending order.
// If reverse is true, the currencies are sorted in descending order.
func (c *CurrencyRates) SortByValue(reverse bool) {
	if reverse {
		sort.Slice(c.Currencies, func(i, j int) bool {
			return c.Currencies[i].Value > c.Currencies[j].Value
		})
	} else {
		sort.Slice(c.Currencies, func(i, j int) bool {
			return c.Currencies[i].Value < c.Currencies[j].Value
		})
	}
}

// If only the filename is provided (without any field names),
// all fields of the CurrencyRates structure will be exported.
func (c *CurrencyRates) ExportSelectedCurrencyRatesToJSON(filename string, fields ...string) {
	var output []map[string]interface{}

	for _, currency := range c.Currencies {
		data := make(map[string]interface{})

		if len(fields) == 0 {
			data["num-code"] = currency.NumCode
			data["char-code"] = currency.CharCode
			data["nominal"] = currency.Nominal
			data["name"] = currency.Name
			data["value"] = currency.Value
			data["vunit-rate"] = currency.VunitRate
		} else {
			for _, field := range fields {
				switch field {
				case "NumCode":
					data["num-code"] = currency.NumCode
				case "CharCode":
					data["char-code"] = currency.CharCode
				case "Nominal":
					data["nominal"] = currency.Nominal
				case "Name":
					data["name"] = currency.Name
				case "Value":
					data["value"] = currency.Value
				case "VunitRate":
					data["vunit-rate"] = currency.VunitRate
				}
			}
		}

		output = append(output, data)
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
