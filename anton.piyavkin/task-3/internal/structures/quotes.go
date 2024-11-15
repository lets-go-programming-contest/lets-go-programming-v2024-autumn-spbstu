package structures

type Quotes struct {
	NumCode   int     `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int     `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float32 `xml:"Value"`
	VunitRate float32 `xml:"VunitRate"`
}
