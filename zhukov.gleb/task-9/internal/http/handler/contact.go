package handler

import (
	"errors"
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

var (
	orderFields = map[string]struct{}{
		"id":         struct{}{},
		"phone":      struct{}{},
		"created_at": struct{}{},
		"updated_at": struct{}{},
	}
)

func (h *ContactHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	orderBy := r.URL.Query().Get("order_by")
	if _, ok := orderFields[orderBy]; !ok {
		WriteJSONServer(w, map[string]string{"message": "Bad request: invalid order_by value"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.GetAll(orderBy)
	if err != nil {
		WriteJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	}

	WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/contacts/{id} - id
	idString := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.GetByID(id)
	switch {
	case errors.Is(err, db.ErrNoContact):
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		WriteJSONServer(w, map[string]string{"message": "internal server error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	value, err := GetJSONFieldsServer(r, "name", "phone")
	switch {
	case errors.Is(err, ErrNoDataJSON):
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		WriteJSONServer(w, map[string]string{"message": "internal server error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	nameReq, nameOk := value["name"].(string)
	phoneReq, phoneOk := value["phone"].(string)
	if !nameOk || !phoneOk {
		WriteJSONServer(w, map[string]string{"message": "Bad request params"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.Add(nameReq, phoneReq)
	switch {
	case errors.Is(err, contact.ErrIncorrectPhone):
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		WriteJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/contacts/{id} - id
	idString := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	}

	value, err := GetJSONFieldsServer(r, "name", "phone")
	switch {
	case errors.Is(err, ErrNoDataJSON):
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		WriteJSONServer(w, map[string]string{"message": "internal server error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	nameReq, nameOk := value["name"].(string)
	phoneReq, phoneOk := value["phone"].(string)
	if !nameOk || !phoneOk {
		WriteJSONServer(w, map[string]string{"message": "Bad request params"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.Update(id, nameReq, phoneReq)
	switch {
	case errors.Is(err, db.ErrNoContact) || errors.Is(err, contact.ErrIncorrectPhone):
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		WriteJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	// http://127.0.0.1:8080/contacts/{id} - id
	idString := strings.Split(r.URL.Path, "/")[2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	}

	err = h.ContactRepo.Delete(id)
	switch {
	case errors.Is(err, db.ErrNoContact):
		WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	case err != nil:
		WriteJSONServer(w, map[string]string{"message": "Internal error: " + err.Error()}, http.StatusInternalServerError)
		return
	default:
	}

	WriteJSONServer(w, map[string]string{"message": "Success delete"}, http.StatusOK)
}
