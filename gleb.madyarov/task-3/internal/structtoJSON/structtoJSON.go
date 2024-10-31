package structtojson

import (
	"encoding/json"
	"fmt"

	xmltostruct "github.com/Madyarov-Gleb/task-3/internal/XMLtoStruct"
)

func StructtoJSON(data *xmltostruct.ValCurs) ([]byte, error) {
	datajsn, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return datajsn, fmt.Errorf("couldn't convert the structure to JSON: %w", err)
	}
	return datajsn, nil
}
