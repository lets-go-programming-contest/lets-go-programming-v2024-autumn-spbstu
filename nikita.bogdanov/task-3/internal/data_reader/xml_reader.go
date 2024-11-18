package data_reader

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/solomonalfred/task-3/internal/schemas"
	"golang.org/x/net/html/charset"
)

func GetValuteCurse(config schemas.ConfigStruct) *schemas.ValuteCurseStructure {
	curseData, err := os.OpenFile(config.Input, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer curseData.Close()
	buf := make([]byte, 1024)
	var data []byte
	for {
		n, state := curseData.Read(buf)
		if state == io.EOF {
			break
		}
		if state != nil {
			panic(state)
		}
		data = append(data, buf[:n]...)
	}
	valCurs := convertToStructure(data)
	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})
	return valCurs
}

func convertToStructure(data []byte) *schemas.ValuteCurseStructure {
	stream := []byte(strings.ReplaceAll(string(data), ",", "."))
	valCurs := new(schemas.ValuteCurseStructure)
	reader := bytes.NewReader(stream)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}
	return valCurs
}
