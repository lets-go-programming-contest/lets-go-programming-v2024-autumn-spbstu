package encoding

import (
	"github.com/nutochk/task-3/internal/structures"
	"os"
	"path/filepath"
)

func OpenOutput(config structures.Config) *os.File {
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
