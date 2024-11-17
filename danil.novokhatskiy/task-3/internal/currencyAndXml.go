package internal

import (
	"bytes"
	"encoding/xml"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type MyFloat float64

func (f *MyFloat) UnmarshalXML(d *xml.Decoder, s xml.StartElement) error {
	var data string
	if err := d.DecodeElement(&data, &s); err != nil {
		return err
	}
	value, err := strconv.ParseFloat(strings.ReplaceAll(string(data), ",", "."), 64)
	if err != nil {
		return err
	}
	*f = MyFloat(value)
	return nil
}

type Currency struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    MyFloat `xml:"Value" json:"value"`
}

type Currencies struct {
	Currencies []Currency `xml:"Valute"`
}

func ParseXml(cur *Currencies, file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&cur)
	if err != nil {
		return err
	}
	return nil
}
