package data_writer

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/solomonalfred/task-3/internal/schemas"
)

func prepareOutput(config schemas.ConfigStruct) *os.File {
	dir := filepath.Dir(config.Output)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
	fileData, err := os.OpenFile(config.Output, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	return fileData
}

func GetJSONReport(config schemas.ConfigStruct, curses *schemas.ValuteCurseStructure) {
	reportData := prepareOutput(config)
	defer reportData.Close()

	var allValuteJSON []schemas.ValuteJSONStructure

	for _, val := range curses.Valute {
		valuteData := schemas.ValuteJSONStructure{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Value:    val.Value,
		}
		allValuteJSON = append(allValuteJSON, valuteData)
	}
	w := bufio.NewWriter(reportData)
	defer w.Flush()
	encd := json.NewEncoder(w)
	encd.SetIndent("", "  ")
	if err := encd.Encode(allValuteJSON); err != nil {
		panic(err)
	}
}
