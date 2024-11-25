package handler

import (
	"encoding/json"
	"net/http"
	"task-9/internal/contact"
)

type ContactHandler struct {
	ContactRepo contact.ContactInface
}

func WriteJSONServer(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (h *ContactHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.PostRepo.GetAll()
	if err != nil {
		h.Log.Errorw("mongo error", "error", err.Error())
		WriteJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	}

	err = h.Tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		h.Log.Errorw("Template execution error", "error", err.Error())
		WriteJSONServer(w, map[string]string{"message": "Template error"}, http.StatusInternalServerError)
		return
	}
}
