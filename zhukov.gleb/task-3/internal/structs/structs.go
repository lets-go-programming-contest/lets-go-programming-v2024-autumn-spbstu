package structs

type Cfg struct {
	InFile  string `yaml:"input-file"`
	OutFile string `yaml:"output-file"`
}

type Currency struct {
	NumCode  int     `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

type ValCurs struct {
	Currencies []Currency `xml:"Valute"`
}
