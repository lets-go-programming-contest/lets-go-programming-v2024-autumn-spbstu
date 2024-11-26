package currency

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
	"bytes"
)

type Currency struct {
	NumCode   int64       `xml:"NumCode" json:"num-code"`
	CharCode  string      `xml:"CharCode" json:"char-code"`
	Nominal   int64       `xml:"Nominal" json:"nominal"`
	Name      string      `xml:"Name" json:"name"`
	Value     CustomFloat `xml:"Value" json:"value"`
	VunitRate CustomFloat `xml:"VunitRate" json:"vunit-rate"`
}

type Currencies struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Date       string     `xml:"Date,attr"`
	Name       string     `xml:"name,attr"`
	Currencies []Currency `xml:"Valute"`
}

type CurrencyJSON struct {
	NumCode   int64       `json:"num-code"`
	CharCode  string      `json:"char-code"`
	Nominal   int64       `json:"nominal"`
	Name      string      `json:"name"`
	Value     CustomFloat `json:"value"`
	VunitRate CustomFloat `json:"vunit-rate"`
}

type CustomFloat float64

func (cf *CustomFloat) UnmarshalText(text []byte) error {
	s := strings.ReplaceAll(string(text), ",", ".")

	value, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*cf = CustomFloat(value)
	return nil
}

func ParseXML(pathToXML string) (*Currencies, error) {
	data, err := os.ReadFile(pathToXML)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	var currencies Currencies
	err = decoder.Decode(&currencies)
	if err != nil {
		return nil, err
	}

	return &currencies, nil
}

func WriteCurrenciesToJSON(filename string, fields []string, pathToXML string) error {
	currencies, err := ParseXML(pathToXML)
	if err != nil {
		return err
	}

	var currenciesJSON []CurrencyJSON
	for _, currency := range currencies.Currencies {
		currenciesJSON = append(currenciesJSON, CurrencyJSON{
			NumCode:   currency.NumCode,
			CharCode:  currency.CharCode,
			Nominal:   currency.Nominal,
			Name:      currency.Name,
			Value:     currency.Value,
			VunitRate: currency.VunitRate,
		})
	}

	if len(fields) > 0 {
		sort.Slice(currenciesJSON, func(i, j int) bool {
			for _, field := range fields {
				switch field {
				case "num-code":
					if currenciesJSON[i].NumCode != currenciesJSON[j].NumCode {
						return currenciesJSON[i].NumCode < currenciesJSON[j].NumCode
					}
				case "char-code":
					if currenciesJSON[i].CharCode != currenciesJSON[j].CharCode {
						return currenciesJSON[i].CharCode < currenciesJSON[j].CharCode
					}
				case "nominal":
					if currenciesJSON[i].Nominal != currenciesJSON[j].Nominal {
						return currenciesJSON[i].Nominal < currenciesJSON[j].Nominal
					}
				case "name":
					if currenciesJSON[i].Name != currenciesJSON[j].Name {
						return currenciesJSON[i].Name < currenciesJSON[j].Name
					}
				case "value":
					if currenciesJSON[i].Value != currenciesJSON[j].Value {
						return currenciesJSON[i].Value < currenciesJSON[j].Value
					}
				case "vunit-rate":
					if currenciesJSON[i].VunitRate != currenciesJSON[j].VunitRate {
						return currenciesJSON[i].VunitRate < currenciesJSON[j].VunitRate
					}
				}
			}
			return false
		})
	}

	jsonData, err := json.MarshalIndent(currenciesJSON, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}