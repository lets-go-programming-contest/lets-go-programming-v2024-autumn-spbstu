package input

import (
	"encoding/xml"
	"github.com/Piyavva/task-3/internal/structures"
	"os"
	"strings"
)

func ReadFile(config structures.Config) structures.Сurrencies {
	file, err := os.ReadFile(config.InputFile)
	if err != nil {
		panic(err)
	}
	file = []byte(strings.ReplaceAll(string(file), ",", "."))
	currencies := structures.Сurrencies{}
	err = xml.Unmarshal(file, &currencies)
	if err != nil {
		panic(err)
	}
	return currencies
}
