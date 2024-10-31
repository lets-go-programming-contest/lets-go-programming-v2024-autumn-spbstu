package writer

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sssidkn/task-3/internal/entities"
)

func WriteFile(path string, outputData *entities.OutputData) error {
	_, err := os.Stat(path)
	if err != nil {
		err = os.MkdirAll(filepath.Dir(path), 0777)
		if err != nil {
			return err
		}
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	extension := filepath.Ext(path)
	var data []byte
	switch extension {
	case ".json":
		data, err = json.MarshalIndent(outputData.Valute, "", "  ")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported file extension: %s", extension)
	}
	err = os.WriteFile(path, data, 0777)
	if err != nil {
		return err
	}
	return nil
}
