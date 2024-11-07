package currencyRates

type Currency struct {
	NumCode   int64   `xml:"NumCode"`
	CharCode  string  `xml:"CharCode"`
	Nominal   int64   `xml:"Nominal"`
	Name      string  `xml:"Name"`
	Value     float64 `xml:"Value"`
	VunitRate float64 `xml:"VunitRate"`
}
