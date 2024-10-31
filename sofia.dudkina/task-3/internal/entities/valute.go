package entities

type OutputData struct {
	Valute []struct {
		NumCode  int     `json:"num_code,omitempty"`
		CharCode string  `json:"char_code,omitempty"`
		Value    float64 `json:"value,omitempty"`
	}
}

type CursData struct {
	Date   string `xml:"Date,attr"`
	Name   string `xml:"name,attr"`
	Valute []struct {
		ID        string `xml:"ID,attr"`
		NumCode   int    `xml:"NumCode"`
		CharCode  string `xml:"CharCode"`
		Nominal   int    `xml:"Nominal"`
		Name      string `xml:"Name"`
		Value     string `xml:"Value"`
		VunitRate string `xml:"VunitRate"`
	}
}

func (outData OutputData) Len() int {
	return len(outData.Valute)
}

func (outData OutputData) Less(i, j int) bool {
	return outData.Valute[i].Value > outData.Valute[j].Value
}

func (outData OutputData) Swap(i, j int) {
	outData.Valute[i], outData.Valute[j] = outData.Valute[j], outData.Valute[i]
}
