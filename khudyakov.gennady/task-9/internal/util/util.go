package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KRYST4L614/task-9/internal/errlib"
)

func WriteJSON(w http.ResponseWriter, jsonStruct interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonStruct)
	if err != nil {
		jsonerr := errlib.GetJSONError(fmt.Errorf("%w: Failed encode", errlib.ErrInternal))
		w.WriteHeader(jsonerr.Error.Code)
		if err = json.NewEncoder(w).Encode(jsonerr); err != nil {
			http.Error(w, "Internal Error", http.StatusInternalServerError)
		}
	}
	w.WriteHeader(code)
}

func WriteJSONError(w http.ResponseWriter, err error) {
	jsonErr := errlib.GetJSONError(err)
	w.WriteHeader(jsonErr.Error.Code)
	WriteJSON(w, jsonErr, jsonErr.Error.Code)
}
