package data_writer

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/solomonalfred/task-3/internal/schemas"
)

func prepareOutput(config schemas.ConfigStruct) (*os.File, error) {
	dir := filepath.Dir(config.Output)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return nil, err
		}
	}
	fileData, err := os.OpenFile(config.Output, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil && !os.IsExist(err) {
		if err != nil {
			return nil, err
		}
	}
	return fileData, nil
}

func GetJSONReport(config schemas.ConfigStruct, curses *schemas.ValuteCurseStructure) error {
	reportData, err := prepareOutput(config)
	if err != nil {
		return err
	}
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
		return err
	}
	return nil
}
