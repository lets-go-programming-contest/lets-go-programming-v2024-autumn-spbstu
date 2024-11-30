package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"task-9/internal/contact"
	"task-9/internal/db"
)

var (
	ErrDecodingJSON = errors.New("error decoding JSON")
	ErrNoDataJSON   = errors.New("error no data JSON")
)

type ContactHandler struct {
	ContactRepo contact.ContactInface
}

func writeJSONServer(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func getJSONFieldsServer(r *http.Request, fields ...string) (map[string]interface{}, error) {
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
	switch {
	case errors.Is(err, db.ErrNoContact):
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		writeJSONServer(w, map[string]string{"message": "internal server error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	writeJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	value, err := getJSONFieldsServer(r, "name", "phone")
	switch {
	case errors.Is(err, ErrNoDataJSON):
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		writeJSONServer(w, map[string]string{"message": "internal server error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	nameReq, nameOk := value["name"].(string)
	phoneReq, phoneOk := value["phone"].(string)
	if !nameOk || !phoneOk {
		writeJSONServer(w, map[string]string{"message": "Bad request params"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.Add(nameReq, phoneReq)
	switch {
	case errors.Is(err, contact.ErrIncorrectPhone):
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		writeJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	writeJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/contacts/{id} - id
	idString := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	}

	value, err := getJSONFieldsServer(r, "name", "phone")
	switch {
	case errors.Is(err, ErrNoDataJSON):
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		writeJSONServer(w, map[string]string{"message": "internal server error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	nameReq, nameOk := value["name"].(string)
	phoneReq, phoneOk := value["phone"].(string)
	if !nameOk || !phoneOk {
		writeJSONServer(w, map[string]string{"message": "Bad request params"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.Update(id, nameReq, phoneReq)
	switch {
	case errors.Is(err, db.ErrNoContact):
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		writeJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	writeJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/contacts/{id} - id
	idString := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	}

	err = h.ContactRepo.Delete(id)
	switch {
	case errors.Is(err, db.ErrNoContact):
		writeJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		writeJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	writeJSONServer(w, map[string]string{"message": "Success delete"}, http.StatusOK)
}
