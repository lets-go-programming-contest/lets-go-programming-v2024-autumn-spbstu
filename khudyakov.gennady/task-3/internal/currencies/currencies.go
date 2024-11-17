package currencies

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"

	"golang.org/x/text/encoding/charmap"
)

type CurrencyValue float64

type Currency struct {
	NumCode  string        `json:"num_code" xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value" xml:"Value"`
}

type Currencies struct {
	Valute []Currency `json:"value" xml:"Valute"`
}

func (c Currencies) Len() int {
	return len(c.Valute)
}

func (c Currencies) Less(i, j int) bool {
	return c.Valute[i].Value > c.Valute[j].Value
}

func (c *Currencies) Swap(i, j int) {
	c.Valute[i], c.Valute[j] = c.Valute[j], c.Valute[i]
}

func (c *Currencies) ParseXML(data []byte) (*Currencies, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "windows-1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		default:
			return nil, fmt.Errorf("unknown charset: %s", charset)
		}
	}

	err := decoder.Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c Currencies) ConvertToJSON() ([]byte, error) {
	data, err := json.Marshal(c.Valute)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (currencyValue *CurrencyValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var data []byte
	err := decoder.DecodeElement(&data, &start)
	if err != nil {
		return err
	}

	data = bytes.ReplaceAll(data, []byte(","), []byte("."))

	result, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}

	*currencyValue = CurrencyValue(result)
	return nil
}
