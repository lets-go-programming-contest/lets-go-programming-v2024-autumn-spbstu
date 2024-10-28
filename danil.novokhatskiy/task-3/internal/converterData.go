package internal

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func Convert(curr *Currencies, file string) error {
	SortCurrencies(curr.Currencies)
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()
	data, err := json.MarshalIndent(curr, "", " ")
	if err != nil {
		return err
	}
	if _, err := out.Write(data); err != nil {
		return err
	}
	return nil
}
