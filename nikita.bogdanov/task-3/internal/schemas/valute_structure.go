package schemas

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type ValuteStructure struct {
	ID        string  `xml:"ID,attr"`
	NumCode   int     `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int     `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float64 `xml:"Value"`
	VunitRate float64 `xml:"VunitRate"`
}

func (v *ValuteStructure) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type Alias ValuteStructure
	aux := &struct {
		Value     string `xml:"Value"`
		VunitRate string `xml:"VunitRate"`
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}
	valueStr := strings.ReplaceAll(aux.Value, ",", ".")
	value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
	if err != nil {
		return fmt.Errorf("error parsing Value: %v", err)
	}
	v.Value = value
	if aux.VunitRate != "" {
		vunitRateStr := strings.ReplaceAll(aux.VunitRate, ",", ".")
		vunitRate, err := strconv.ParseFloat(strings.TrimSpace(vunitRateStr), 64)
		if err != nil {
			return fmt.Errorf("error parsing VunitRate: %v", err)
		}
		v.VunitRate = vunitRate
	} else {
		v.VunitRate = 0.0
	}
	v.ID = aux.ID
	v.NumCode = aux.NumCode
	v.CharCode = aux.CharCode
	v.Nominal = aux.Nominal
	v.Name = aux.Name
	return nil
}
