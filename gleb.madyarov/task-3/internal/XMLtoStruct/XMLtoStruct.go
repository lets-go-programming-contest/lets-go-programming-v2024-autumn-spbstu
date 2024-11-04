package xmltostruct

import (
	"bytes"
	"encoding/xml"
	"fmt"
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

func XMLtoStruct(data *ValCurs, config read.Config) (ValCurs, error) {
	file, err := os.ReadFile(config.Input)
	if err != nil {
		return *data, fmt.Errorf("couldn't read the input file: %w", err)
	}
	file = []byte(strings.ReplaceAll(string(file), ",", "."))
	r := bytes.NewReader([]byte(file))
	d := xml.NewDecoder(r)
	d.CharsetReader = charset.NewReaderLabel
	err = d.Decode(&data)
	if err != nil {
		return *data, fmt.Errorf("it was not possible to convert XML into a structure: %w", err)
	}
	return *data, nil
}
