package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	myErr "github.com/EmptyInsid/task-9/internal/errors"
	"github.com/EmptyInsid/task-9/internal/models"
)

type dbService interface {
	GetContacts() ([]models.Contact, error)
	GetContact(id int) (*models.Contact, error)
	CreateContact(contact models.Contact) (int, error)
	UpdateContact(contact models.Contact) (*models.Contact, error)
	DeleteContact(id int) error
}

type handler struct {
	service dbService
}

func NewHandler(service dbService, router *mux.Router) *mux.Router {
	h := handler{
		service: service,
	}

	router.HandleFunc("/contacts", h.getContacts).Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/contacts/{id}", h.getContact).Methods(http.MethodOptions, http.MethodGet)
	router.HandleFunc("/contacts", h.createContact).Methods(http.MethodOptions, http.MethodPost)
	router.HandleFunc("/contacts/{id}", h.updateContact).Methods(http.MethodOptions, http.MethodPut)
	router.HandleFunc("/contacts/{id}", h.deleteContact).Methods(http.MethodOptions, http.MethodDelete)

	return router
}

func (h *handler) getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contacts, err := h.service.GetContacts()
	if err != nil {
		http.Error(w, err.Error(), getStatusCode(errorsGetMap, err))

		return
	}

	if err = json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, myErr.ErrEncodeJSON.Error(), getStatusCode(errorsGetMap, myErr.ErrEncodeJSON))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, myErr.ErrNoContact.Error(), getStatusCode(errorsGetMap, myErr.ErrNoContact))

		return
	}

	contact, err := h.service.GetContact(id)
	if err != nil {
		http.Error(w, err.Error(), getStatusCode(errorsGetMap, err))

		return
	}

	if err = json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, myErr.ErrEncodeJSON.Error(), getStatusCode(errorsGetMap, myErr.ErrEncodeJSON))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, myErr.ErrDecodeJSON.Error(), getStatusCode(errorsCreateMap, myErr.ErrDecodeJSON))

		return
	} else if isEmpty(contact.Name) || isEmpty(contact.Phone) {
		http.Error(w, myErr.ErrEmptyData.Error(), getStatusCode(errorsCreateMap, myErr.ErrEmptyData))

		return
	} else if !isValidPhone(contact.Phone) {
		http.Error(w, myErr.ErrWrongPhoneFormat.Error(), getStatusCode(errorsCreateMap, myErr.ErrWrongPhoneFormat))

		return
	}

	id, err := h.service.CreateContact(contact)
	if err != nil {
		http.Error(w, err.Error(), getStatusCode(errorsCreateMap, err))

		return
	}

	contact.ID = id
	if err = json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, myErr.ErrEncodeJSON.Error(), getStatusCode(errorsCreateMap, myErr.ErrEncodeJSON))

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handler) updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, myErr.ErrNoContact.Error(), getStatusCode(errorsUpdMap, myErr.ErrNoContact))

		return
	}

	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, myErr.ErrDecodeJSON.Error(), getStatusCode(errorsCreateMap, myErr.ErrDecodeJSON))

		return
	} else if isEmpty(contact.Name) || isEmpty(contact.Phone) {
		http.Error(w, myErr.ErrEmptyData.Error(), getStatusCode(errorsUpdMap, myErr.ErrEmptyData))

		return
	} else if !isEmpty(contact.Phone) && !isValidPhone(contact.Phone) {
		http.Error(w, myErr.ErrWrongPhoneFormat.Error(), getStatusCode(errorsUpdMap, myErr.ErrWrongPhoneFormat))

		return
	}

	contact.ID = id

	newContact, err := h.service.UpdateContact(contact)
	if err != nil {
		http.Error(w, err.Error(), getStatusCode(errorsUpdMap, err))

		return
	}

	if err := json.NewEncoder(w).Encode(newContact); err != nil {
		http.Error(w, myErr.ErrEncodeJSON.Error(), getStatusCode(errorsUpdMap, myErr.ErrEncodeJSON))

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handler) deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, myErr.ErrNoContact.Error(), getStatusCode(errorsDeleteMap, myErr.ErrNoContact))

		return
	}

	if err := h.service.DeleteContact(id); err != nil {
		http.Error(w, err.Error(), getStatusCode(errorsDeleteMap, err))

		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"contact": "deleted"})
	if err != nil {
		http.Error(w, myErr.ErrEncodeJSON.Error(), getStatusCode(errorsDeleteMap, myErr.ErrEncodeJSON))

		return
	}

	w.WriteHeader(http.StatusOK)
}
