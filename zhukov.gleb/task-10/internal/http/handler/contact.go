package handler

import (
	"errors"
	"net/http"
	"strconv"

	"task-10/internal/contact"
	"task-10/internal/db"
	"task-10/internal/http/encode"

	"github.com/gorilla/mux"
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
		encode.WriteJSONServer(w, map[string]string{"message": "invalid order_by value"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.GetAll(orderBy)
	if err != nil {
		encode.WriteJSONServer(w, []interface{}{}, http.StatusInternalServerError)
		return
	}

	encode.WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		encode.WriteJSONServer(w, map[string]string{"message": "bad integer arg"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.GetByID(id)
	reqArgs := map[string]interface{}{
		"reqArgs": map[string]interface{}{
			"id": id,
		},
	}

	switch {
	case err == nil:
	case errors.Is(err, db.ErrNoContact):
		reqArgs["message"] = db.ErrNoContact.Error()
		encode.WriteJSONServer(w, reqArgs, http.StatusUnprocessableEntity)
		return
	default:
		encode.WriteJSONServer(w, reqArgs, http.StatusInternalServerError)
		return
	}

	encode.WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	value, err := encode.GetJSONFieldsServer(r, "name", "phone")
	switch {
	case err == nil:
	case errors.Is(err, encode.ErrNoDataJSON):
		encode.WriteJSONServer(w, map[string]string{"message": encode.ErrNoDataJSON.Error()}, http.StatusBadRequest)
		return
	default:
		encode.WriteJSONServer(w, []interface{}{}, http.StatusInternalServerError)
		return
	}

	nameReq, nameOk := value["name"].(string)
	phoneReq, phoneOk := value["phone"].(string)
	if !nameOk || !phoneOk {
		encode.WriteJSONServer(w, map[string]string{"message": "Bad request params"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.Add(nameReq, phoneReq)
	reqArgs := map[string]interface{}{
		"reqArgs": map[string]interface{}{
			"name":  nameReq,
			"phone": phoneReq,
		},
	}

	switch {
	case err == nil:
	case errors.Is(err, contact.ErrIncorrectPhone):
		reqArgs["message"] = contact.ErrIncorrectPhone.Error()
		encode.WriteJSONServer(w, reqArgs, http.StatusBadRequest)
		return
	default:
		encode.WriteJSONServer(w, reqArgs, http.StatusInternalServerError)
		return
	}

	encode.WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		encode.WriteJSONServer(w, map[string]string{"message": "bad integer arg"}, http.StatusBadRequest)
		return
	}

	value, err := encode.GetJSONFieldsServer(r, "name", "phone")
	switch {
	case err == nil:
	case errors.Is(err, encode.ErrNoDataJSON):
		encode.WriteJSONServer(w, map[string]string{"message": "Bad request: " + err.Error()}, http.StatusBadRequest)
		return
	default:
		encode.WriteJSONServer(w, []interface{}{}, http.StatusInternalServerError)
		return
	}

	nameReq, nameOk := value["name"].(string)
	phoneReq, phoneOk := value["phone"].(string)
	if !nameOk || !phoneOk {
		encode.WriteJSONServer(w, map[string]string{"message": "Bad request params"}, http.StatusBadRequest)
		return
	}

	data, err := h.ContactRepo.Update(id, nameReq, phoneReq)
	reqArgs := map[string]interface{}{
		"reqArgs": map[string]interface{}{
			"id":    id,
			"name":  nameReq,
			"phone": phoneReq,
		},
	}

	switch {
	case err == nil:
	case errors.Is(err, contact.ErrIncorrectPhone):
		reqArgs["message"] = contact.ErrIncorrectPhone.Error()
		encode.WriteJSONServer(w, reqArgs, http.StatusBadRequest)
		return
	case errors.Is(err, db.ErrNoContact):
		reqArgs["message"] = db.ErrNoContact.Error()
		encode.WriteJSONServer(w, reqArgs, http.StatusUnprocessableEntity)
		return
	default:
		encode.WriteJSONServer(w, reqArgs, http.StatusInternalServerError)
		return
	}

	encode.WriteJSONServer(w, data, http.StatusOK)
}

func (h *ContactHandler) DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		encode.WriteJSONServer(w, map[string]string{"message": "bad integer arg"}, http.StatusBadRequest)
		return
	}

	err = h.ContactRepo.Delete(id)
	reqArgs := map[string]interface{}{
		"reqArgs": map[string]interface{}{
			"id": id,
		},
	}

	switch {
	case err == nil:
	case errors.Is(err, db.ErrNoContact):
		reqArgs["message"] = db.ErrNoContact.Error()
		encode.WriteJSONServer(w, reqArgs, http.StatusUnprocessableEntity)
		return
	default:
		encode.WriteJSONServer(w, reqArgs, http.StatusInternalServerError)
		return
	}

	reqArgs["message"] = "Success delete"
	encode.WriteJSONServer(w, reqArgs, http.StatusOK)
}
