package valuteProcessing

import (
	"encoding/json"
	"fmt"
	"task-3/internal/valuteStrukts"
)

func ParseToJSON(data *valuteStrukts.ValuteRate) ([]byte, error) {
	dataJSON, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return nil, fmt.Errorf("failed to convert to JSON: %w", err)
	}
	return dataJSON, nil
}
