package structs

type Valute struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float32 `xml:"Value" json:"value"`
	//	ID        int    `xml:"ID,attr"`
	//	VunitRate int    `xml:"VunitRate"`
	// Nominal   int    `xml:"Nominal"`
	// Name      string `xml:"Name"`
}
