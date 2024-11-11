package src

import (
	"bytes"
	"encoding/xml"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/nikita.bogdanov/task-3/internal/shemas"
)

func GetData(config structures.ConfigStruct) *structures.ValuteCurse {
	inputFile, err := os.OpenFile(config.Input, os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	buf := make([]byte, 512)
	var data []byte
	for {
		n, err := inputFile.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		data = append(data, buf[:n]...)
	}

	data = []byte(strings.ReplaceAll(string(data), ",", "."))

	valCurs := new(structures.ValuteCurse)
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&valCurs)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valute, func(i, j int) bool {
		return valCurs.Valute[i].Value < valCurs.Valute[j].Value
	})
	return valCurs
}
