package encoding

import (
	"os"
	"path/filepath"
	"bufio"
	"encoding/json"

	"github.com/solomonalfred/task-3/internal/shemas"
)

func openOutput(config structures.ConfigStruct) *os.File {
	dir := filepath.Dir(config.Output)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	outputFile, err := os.OpenFile(config.Output, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	return outputFile
}

func GetJson(config structures.ConfigStruct, valCurs *structures.ValuteCurse) {
	outputFile := openOutput(config)
	defer outputFile.Close()
	var vij []structures.ValuteJSON

	for _, valute := range valCurs.Valute {
		jsonVal := structures.ValuteJSON{
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
