package strtructs

import (
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
