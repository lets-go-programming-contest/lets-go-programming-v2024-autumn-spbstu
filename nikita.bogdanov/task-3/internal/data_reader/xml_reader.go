package data_reader

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strings"

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
	decoder := charmap.Windows1251.NewDecoder()
	reader := transform.NewReader(file, decoder)
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// data = []byte(data)
	data = []byte(strings.ReplaceAll(string(data), ",", "."))
	data = fixXMLPrologEncoding(data)
	valCurs := new(schemas.ValuteCurseStructure)
	err = xml.Unmarshal(data, valCurs)
	if err != nil {
		return nil, err
	}
	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})
	return valCurs, nil
}

func fixXMLPrologEncoding(data []byte) []byte {
	oldEncoding := []byte(`encoding="windows-1251"`)
	newEncoding := []byte(`encoding="UTF-8"`)
	data = bytes.Replace(data, oldEncoding, newEncoding, 1)
	return data
}
