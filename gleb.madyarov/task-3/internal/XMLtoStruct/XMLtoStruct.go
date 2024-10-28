package xmltostruct

import (
	"bytes"
	"encoding/xml"
	"os"
	"strings"

	"github.com/Madyarov-Gleb/task-3/internal/read"
	"golang.org/x/net/html/charset"
)

type ValCurs struct {
	ValCurs []Valute `xml:"Valute" json:"valute"`
}

type Valute struct {
	NumCode  string  `xml:"NumCode" json:"num_code"`
	CharCode string  `xml:"CharCode" json:"char_code"`
	Value    float64 `xml:"Value" json:"value"`
}

func XMLtoStruct(data *ValCurs, config read.Config) ValCurs {
	file, err := os.ReadFile(config.Input)
	if err != nil {
		panic("couldn't read the input file")
	}
	file = []byte(strings.ReplaceAll(string(file), ",", "."))
	r := bytes.NewReader([]byte(file))
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	err = d.Decode(&data)
	if err != nil {
		panic("it was not possible to convert XML into a structure")
	}
	return *data
}
