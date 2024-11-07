package currencyRates

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type CurrencyRates struct {
	XMLName    xml.Name   `xml:"ValCurs"` // Name of the root element
	Date       string     `xml:"Date,attr"`
	Name       string     `xml:"name,attr"` // Name of the currency market
	Currencies []Currency `xml:"Valute"`
}

func (c *CurrencyRates) ParseXML(pathToXML string) error {
	file, err := os.ReadFile(pathToXML)
	if err != nil {
		return err
	}

	modifiedContent := strings.ReplaceAll(string(file), ",", ".")

	err = xml.Unmarshal([]byte(modifiedContent), c)
	if err != nil {
		return err
	}

	return nil
}

// If reverse is false, the currencies are sorted in ascending order by default.
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

// ExportSelectedCurrencyRatesToJSON exports the selected currency rates to a JSON file.
// If only the filename is provided (without any field names),
// all fields of the CurrencyRates structure will be exported.
func (c *CurrencyRates) ExportSelectedCurrencyRatesToJSON(filename string, fields ...string) error {
	var output []map[string]interface{}

	for _, currency := range c.Currencies {
		data := make(map[string]interface{})

		if len(fields) == 0 {
			data["num-сode"] = currency.NumCode
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
		return err
	}

	err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
