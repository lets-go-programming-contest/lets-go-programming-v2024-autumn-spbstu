package read

import (
	"encoding/xml"
	"os"
	"strings"

	"github.com/artem6554/task-3/structs"
)

func ParseXML(inputFile string) structs.ValCurs {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	data = []byte(strings.ReplaceAll(string(data), ",", "."))

	var currencies structs.ValCurs
	err = xml.Unmarshal(data, &currencies)
	if err != nil {
		panic(err)
	}
	return currencies
}
