package structtojson

import (
	"encoding/json"

	xmltostruct "github.com/Madyarov-Gleb/task-3/internal/XMLtoStruct"
)

func StructtoJSON(data *xmltostruct.ValCurs) []byte {
	datajsn, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		panic("couldn't convert the structure to JSON")
	}
	return datajsn
}
