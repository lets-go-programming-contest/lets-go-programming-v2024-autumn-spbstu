package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSONServer(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal error: "+err.Error(), http.StatusInternalServerError)
	}
}

func GetJSONFieldsServer(r *http.Request, fields ...string) (map[string]interface{}, error) {
	var requestData map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDecodingJSON, err)
	}

	result := make(map[string]interface{})

	for _, field := range fields {
		value, ok := requestData[field]
		if !ok {
			return nil, fmt.Errorf("%w: field '%s' not found", ErrNoDataJSON, field)
		}
		result[field] = value
	}

	return result, nil
}
