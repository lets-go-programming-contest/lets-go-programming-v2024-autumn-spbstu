package structs

type ValCurs struct {
	// XMLName xml.Name `xml:"ValCurs"`
	// Date    string   `xml:"Date",attr`
	// Name    string   `xml:"name",attr`
	Valutes []Valute `xml:"Valute"`
}
