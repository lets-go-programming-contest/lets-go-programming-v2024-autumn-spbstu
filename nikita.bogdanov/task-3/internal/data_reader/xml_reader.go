package data_reader

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/solomonalfred/task-3/internal/schemas"
)

func GetValuteCurse(config schemas.ConfigStruct) (*schemas.ValuteCurseStructure, error) {
	curseData, err := os.OpenFile(config.Input, os.O_RDONLY, 0777)
	if err != nil {
		return nil, err
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
			return nil, state
		}
		data = append(data, buf[:n]...)
	}
	valCurs, err := convertToStructure(data)
	if err != nil {
		return nil, err
	}
	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})
	return valCurs, nil
}

func convertToStructure(data []byte) (*schemas.ValuteCurseStructure, error) {
	stream := []byte(strings.ReplaceAll(string(data), ",", ".")) //Todo: Unmarshal
	valCurs := new(schemas.ValuteCurseStructure)
	reader := bytes.NewReader(stream)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&valCurs)
	if err != nil {
		return nil, err
	}
	return valCurs, nil
}
