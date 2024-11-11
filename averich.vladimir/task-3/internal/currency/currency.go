package currency

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
	"strings"
	"golang.org/x/net/html/charset"
	"bytes"
	"reflect"
)

type Currency struct {
	NumCode   int64   `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int64   `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float64 `xml:"Value"`
	VunitRate float64 `xml:"VunitRate"`
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
	var currenciesJSON []map[string]interface{}
	for _, currency := range currencies.Currencies {
		currencyMap := make(map[string]interface{})
		v := reflect.ValueOf(currency)
		t := v.Type()

		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				continue
			}
			if len(fields) == 0 || contains(fields, jsonTag) {
				currencyMap[jsonTag] = v.Field(i).Interface()
			}
		}

		currenciesJSON = append(currenciesJSON, currencyMap)
	}

	data, err := json.MarshalIndent(currenciesJSON, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}