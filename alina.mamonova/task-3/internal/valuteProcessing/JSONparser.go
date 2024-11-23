package valuteProcessing

import (
	"encoding/json"
	"fmt"
	"github.com/hahapathetic/task-3/internal/valuteStructs"
)

func ParseToJSON(data *valuteStructs.ValuteRate) ([]byte, error) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return dataJSON, fmt.Errorf("failed to convert to JSON: %w", err)
	}
	return dataJSON, nil
}
