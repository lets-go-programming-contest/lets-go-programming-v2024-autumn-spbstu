package write

import (
	"encoding/json"
	"os"
	"path/filepath"

	structs "task-3/internal/structs"
)

func WriteToJson(currencies structs.ValCurs, outFile string) {
	dir := filepath.Dir(outFile)
	os.MkdirAll(dir, os.ModePerm)
	file, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	data, err := json.MarshalIndent(currencies.Valutes, "", "\t")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
