package input

import (
	"encoding/xml"
	"github.com/Piyavva/task-3/internal/structures"
	"golang.org/x/net/html/charset"
	"os"
	"strings"
)

func ReadFile(config structures.Config) structures.Сurrencies {
	file, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}
	file = []byte(strings.ReplaceAll(string(file), ",", "."))
	currencies := structures.Сurrencies{}
	dec := xml.NewDecoder(strings.NewReader(string(file)))
	dec.CharsetReader = charset.NewReaderLabel
	err = dec.Decode(&currencies)
	//err = xml.Unmarshal(file, &currencies)
	if err != nil {
		panic(err)
	}
	return currencies
}
