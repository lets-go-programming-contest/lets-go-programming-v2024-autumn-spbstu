package shemas


type Valute struct {
	ID        string  `xml:"ID,attr"`
	NumCode   int     `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int     `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float64 `xml:"Value"`
	VunitRate float64 `xml:"VunitRate"`
}
