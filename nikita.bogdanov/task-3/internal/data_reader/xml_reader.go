package data_reader

import (
	"encoding/xml"
	"os"
	"sort"

	"github.com/solomonalfred/task-3/internal/schemas"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func GetValuteCurse(config schemas.ConfigStruct) (*schemas.ValuteCurseStructure, error) {
	file, err := os.Open(config.Input)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := xml.NewDecoder(transform.NewReader(file, charmap.Windows1251.NewDecoder()))
	valCurs := new(schemas.ValuteCurseStructure)
	err = decoder.Decode(valCurs)
	if err != nil {
		return nil, err
	}
	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})
	return valCurs, nil
}
