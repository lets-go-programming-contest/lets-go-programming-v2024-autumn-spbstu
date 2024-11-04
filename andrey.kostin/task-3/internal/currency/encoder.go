package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	errProc "github.com/IDevFrye/task-3/internal/errors"
)

func SaveAsJSON(data []Currency, filePath string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrJSONMarshal, err.Error())
	}

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrDirectoryCreation, err.Error())
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrFileCreation, err.Error())
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("%w: %s", errProc.ErrDataWriting, err.Error())
	}

	return nil
}
