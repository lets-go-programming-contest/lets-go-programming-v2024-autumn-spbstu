package entities

type ValCurs struct {
	Valute []struct {
		NumCode  int     `xml:"NumCode" json:"num_code,omitempty"`
		CharCode string  `xml:"CharCode" json:"char_code,omitempty"`
		Value    float64 `xml:"Value" json:"value,omitempty"`
	}
}

func (outData ValCurs) Len() int {
	return len(outData.Valute)
}

func (outData ValCurs) Less(i, j int) bool {
	return outData.Valute[i].Value > outData.Valute[j].Value
}

func (outData ValCurs) Swap(i, j int) {
	outData.Valute[i], outData.Valute[j] = outData.Valute[j], outData.Valute[i]
}
