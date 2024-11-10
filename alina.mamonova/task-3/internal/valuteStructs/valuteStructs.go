package valuteStructs

type ValuteRate struct {
	ValuteRate []Valute `xml:"Valute" json:"valute"`
}

type Valute struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}
