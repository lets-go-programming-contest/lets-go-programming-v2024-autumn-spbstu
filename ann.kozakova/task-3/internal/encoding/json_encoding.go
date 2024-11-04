package encoding

import (
	"bufio"
	"encoding/json"
	"github.com/nutochk/task-3/internal/structures"
)

func JsonEncoding(config structures.Config, valCurs *structures.ValCurs) {
	outputFile := OpenOutput(config)
	defer outputFile.Close()
	var vij []structures.ValuteInJSON

	for _, valute := range valCurs.Valute {
		jsonVal := structures.ValuteInJSON{
			NumCode:  valute.NumCode,
			CharCode: valute.CharCode,
			Value:    valute.Value,
		}
		vij = append(vij, jsonVal)
	}

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(vij); err != nil {
		panic(err)
	}
}
