package strtructs

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
	"strings"
)

type CurrencyRates struct {
	XMLName    xml.Name   `xml:"ValCurs"`   // Name of the root element
	Date       string     `xml:"Date,attr"` // Date
	Name       string     `xml:"name,attr"` // Name of the currency market
	Currencies []Currency `xml:"Valute"`    // Array of currencies
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

func (c *CurrencyRates) ExportSelectedCurrencyRatesToJSON(filename string, fields ...string) error {
	var output []map[string]interface{}

	for _, currency := range c.Currencies {
		data := make(map[string]interface{})

		if len(fields) == 0 {
			data["NumCode"] = currency.NumCode
			data["CharCode"] = currency.CharCode
			data["Nominal"] = currency.Nominal
			data["Name"] = currency.Name
			data["Value"] = currency.Value
			data["VunitRate"] = currency.VunitRate
		} else {
			for _, field := range fields {
				switch field {
				case "NumCode":
					data["NumCode"] = currency.NumCode
				case "CharCode":
					data["CharCode"] = currency.CharCode
				case "Nominal":
					data["Nominal"] = currency.Nominal
				case "Name":
					data["Name"] = currency.Name
				case "Value":
					data["Value"] = currency.Value
				case "VunitRate":
					data["VunitRate"] = currency.VunitRate
				}
			}
		}

		output = append(output, data)
	}

	jsonData, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
