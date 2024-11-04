package entities

type CursData struct {
	ValCurs struct {
		Date   string `xml:"Date,attr"`
		Name   string `xml:"name,attr"`
		Valute []struct {
			ID        string  `xml:"ID,attr" json:"-"`
			NumCode   int     `xml:"NumCode" json:"num_code,omitempty"`
			CharCode  string  `xml:"CharCode" json:"char_code,omitempty"`
			Nominal   int     `xml:"Nominal" json:"-"`
			Name      string  `xml:"Name" json:"-"`
			Value     float64 `xml:"Value" json:"value,omitempty"`
			VunitRate float64 `xml:"VunitRate" json:"-"`
		} `xml:"Valute" json:"Valute,omitempty"`
	} `xml:"ValCurs" json:"ValCurs,omitempty"`
}

func (cursData CursData) Len() int {
	return len(cursData.ValCurs.Valute)
}

func (cursData CursData) Less(i, j int) bool {
	return cursData.ValCurs.Valute[i].Value > cursData.ValCurs.Valute[j].Value
}

func (cursData CursData) Swap(i, j int) {
	cursData.ValCurs.Valute[i], cursData.ValCurs.Valute[j] = cursData.ValCurs.Valute[j], cursData.ValCurs.Valute[i]
}
