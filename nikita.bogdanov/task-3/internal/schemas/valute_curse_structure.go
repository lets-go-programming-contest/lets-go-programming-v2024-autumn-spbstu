package schemas

import "encoding/xml"

type ValuteCurseStructure struct {
	XMLName xml.Name          `xml:"ValCurs"`
	Date    string            `xml:"Date,attr"`
	Name    string            `xml:"name,attr"`
	Valute  []ValuteStructure `xml:"Valute"`
}
