package write

import (
	"encoding/json"
	"os"
	"path/filepath"

	errors "task-3/internal/errorsExt"
	structs "task-3/internal/readFile"
)

func WriteToJson(currencies structs.ValCurs, outFile string) error {
	dir := filepath.Dir(outFile)
	err := os.MkdirAll(dir, os.ModePerm) // обработка ошибки
	if err != nil {
		return errors.ErrorLocate(err)
	}

	file, err := os.Create(outFile)
	if err != nil {
		return errors.ErrorLocate(err)
	}

	data, err := json.MarshalIndent(currencies.Valutes, "", "\t")
	if err != nil {
		return errors.ErrorLocate(err)
	}

	_, err = file.Write(data)
	if err != nil {
		return errors.ErrorLocate(err)
	}

	err = file.Close()
	if err != nil {
		return errors.ErrorLocate(err)
	}

	return nil
}
