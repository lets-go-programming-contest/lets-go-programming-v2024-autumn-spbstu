package structures

import "encoding/xml"

type Сurrencies struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Quotes  []Quotes `xml:"Valute"`
}
