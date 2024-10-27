package entities

type ValCurs struct {
	Date   string `xml:"Date,attr" json:"Date,omitempty"`
	Name   string `xml:"name,attr" json:"Name,omitempty"`
	Valute []struct {
		ID        string `xml:"ID,attr" json:"ID,omitempty"`
		NumCode   string `xml:"NumCode" json:"NumCode,omitempty"`
		CharCode  string `xml:"CharCode" json:"CharCode,omitempty"`
		Nominal   string `xml:"Nominal" json:"Nominal,omitempty"`
		Name      string `xml:"Name" json:"Name,omitempty"`
		Value     string `xml:"Value" json:"Value,omitempty"`
		VunitRate string `xml:"VunitRate" json:"VunitRate,omitempty"`
	} `xml:"Valute" json:"Valute,omitempty"`
}

func (valCurs ValCurs) Len() int {
	return len(valCurs.Valute)
}

func (valCurs ValCurs) Less(i, j int) bool {
	return valCurs.Valute[i].Value < valCurs.Valute[j].Value
}

func (valCurs ValCurs) Swap(i, j int) {
	valCurs.Valute[i], valCurs.Valute[j] = valCurs.Valute[j], valCurs.Valute[i]
}
