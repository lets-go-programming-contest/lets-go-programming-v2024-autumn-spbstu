package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KRYST4L614/task-9/internal/domain"
	"github.com/KRYST4L614/task-9/internal/errlib"
	"github.com/KRYST4L614/task-9/internal/service"
	"github.com/KRYST4L614/task-9/internal/util"
	"github.com/gorilla/mux"
)

type ContactHandler struct {
	service service.ContractServiceIface
}

func NewContactHandler(service service.ContractServiceIface) *ContactHandler {
	return &ContactHandler{
		service: service,
	}
}

func (handler *ContactHandler) GetContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		util.WriteJSONError(w, fmt.Errorf("%w: Invalid path variable id", errlib.ErrBadRequest))
		return
	}

	foundContact, err := handler.service.GetContact(r.Context(), id)
	if err != nil {
		util.WriteJSONError(w, err)
		return
	}

	util.WriteJSON(w, foundContact, http.StatusOK)
}

func (handler *ContactHandler) GetAllContacts(w http.ResponseWriter, r *http.Request) {
	foundContacts, err := handler.service.GetAllContacts(r.Context())
	if err != nil {
		util.WriteJSONError(w, err)
		return
	}

	util.WriteJSON(w, foundContacts, http.StatusOK)
}

func (handler *ContactHandler) AddContact(w http.ResponseWriter, r *http.Request) {
	var contact domain.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		util.WriteJSONError(w, fmt.Errorf("%w: Invalid input", errlib.ErrBadRequest))
		return
	}

	createdContact, err := handler.service.AddContact(r.Context(), contact)
	if err != nil {
		util.WriteJSONError(w, err)
		return
	}

	util.WriteJSON(w, createdContact, http.StatusCreated)
}

func (handler *ContactHandler) UpdateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		util.WriteJSONError(w, fmt.Errorf("%w: Invalid path variable id", errlib.ErrBadRequest))
		return
	}

	var contact domain.Contact
	err = json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		util.WriteJSONError(w, fmt.Errorf("%w: Invalid input", errlib.ErrBadRequest))
		return
	}

	contact.Id = id

	updatedContact, err := handler.service.UpdateContact(r.Context(), contact)
	if err != nil {
		util.WriteJSONError(w, err)
		return
	}

	util.WriteJSON(w, updatedContact, http.StatusOK)
}

func (handler *ContactHandler) DeleteContactById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		util.WriteJSONError(w, fmt.Errorf("%w: Invalid path variable id", errlib.ErrBadRequest))
		return
	}

	err = handler.service.DeleteContactById(r.Context(), id)
	if err != nil {
		util.WriteJSONError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
