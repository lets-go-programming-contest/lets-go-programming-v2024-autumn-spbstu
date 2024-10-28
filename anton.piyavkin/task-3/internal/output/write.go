package output

import (
	"encoding/json"
	"github.com/Piyavva/task-3/internal/structures"
	"os"
	"path/filepath"
)

func Write(val structures.Сurrencies, cfg structures.Config) {
	directory := filepath.Dir(cfg.OutputFile)
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		panic(err)
	}
	file, err := os.OpenFile(cfg.OutputFile, os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	var valutes []structures.OutputVal
	for _, valute := range val.Quotes {
		valutes = append(valutes, structures.OutputVal{valute.NumCode, valute.CharCode, valute.Value})
	}
	data, err := json.Marshal(valutes)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(file.Name(), data, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
