package xmlStruct

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/EmptyInsid/task-3/internal/errorUtils"
)

type float64Comma float64

func (f *float64Comma) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var valueStr string
	if err := d.DecodeElement(&valueStr, &start); err != nil {
		return errorUtils.ErrorWithLocation(err)
	}

	if valueStr == "" {
		*f = 0
		return nil
	}

	valueStr = strings.Replace(valueStr, ",", ".", 1)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return errorUtils.ErrorWithLocation(err)
	}
	*f = float64Comma(value)
	return nil
}

type DataCursStruct struct {
	ValCurs ValCursStruct `xml:"ValCurs" json:"ValCurs,omitempty"`
}

type ValCursStruct struct {
	Date   string   `xml:"Date,attr" json:"Date"`
	Name   string   `xml:"Name,attr" json:"Name"`
	Valute []Valute `xml:"Valute" json:"Valute,omitempty"`
}

type Valute struct {
	ID        string       `xml:"ID,attr" json:"-"`
	NumCode   int          `xml:"NumCode" json:"NumCode,omitempty"`
	CharCode  string       `xml:"CharCode" json:"CharCode,omitempty"`
	Nominal   int          `xml:"Nominal" json:"-"`
	Name      string       `xml:"Name" json:"-"`
	Value     float64Comma `xml:"Value" json:"Value,omitempty"`
	VunitRate float64Comma `xml:"VunitRate" json:"-"`
}

func (data DataCursStruct) Len() int {
	return len(data.ValCurs.Valute)
}

func (data DataCursStruct) Less(i, j int) bool {
	return data.ValCurs.Valute[i].Value > data.ValCurs.Valute[j].Value
}

func (data DataCursStruct) Swap(i, j int) {
	data.ValCurs.Valute[i], data.ValCurs.Valute[j] = data.ValCurs.Valute[j], data.ValCurs.Valute[i]
}
