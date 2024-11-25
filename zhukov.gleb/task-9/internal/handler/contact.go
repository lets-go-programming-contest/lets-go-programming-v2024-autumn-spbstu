package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"task-9/internal/contact"
)

type ContactHandler struct {
	ContactRepo contact.ContactInface
}

func (h *ContactHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	data, err := h.ContactRepo.GetAll()
	if err != nil {
		writeJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/contacts/{id} - id
	idString := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.GetByID(id)
	if err != nil {
		writeJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	}

	writeJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}

func writeJSONServer(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
