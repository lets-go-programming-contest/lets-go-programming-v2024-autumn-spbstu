package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/EmptyInsid/task-9/internal/models"
	"github.com/gorilla/mux"
)

type dbService interface {
	GetContacts() ([]models.Contact, error)
	GetContact(id int) (*models.Contact, error)
	CreateContact(contact models.Contact) (int, error)
	UpdateContact(contact models.Contact) error
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
	router.HandleFunc("/contacts/{id:[0-9]+}", h.updateContact).Methods(http.MethodOptions, http.MethodPut)
	router.HandleFunc("/contacts/{id:[0-9]+}", h.deleteContact).Methods(http.MethodOptions, http.MethodDelete)

	return router
}

func (h *handler) getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	contacts, err := h.service.GetContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err = json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, "Error while encoding contacts", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *handler) getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := h.service.GetContact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, "Error while encoding contact", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *handler) createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Error while decode contact for creating", http.StatusBadRequest)
		return
	} else if isEmpty(contact.Name) || isEmpty(contact.Phone) || !isValidPhone(contact.Phone) {
		http.Error(w, "Incorrect name or phone number: they shouldn't be empty and number should be like: +7 (800) 555-35-55", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateContact(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	contact.Id = id
	if err = json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, "Error while encoding contact", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (h *handler) updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Error while decode contact for creating", http.StatusBadRequest)
		return
	} else if isEmpty(contact.Name) && isEmpty(contact.Phone) {
		http.Error(w, "Name or phone number should be not empty", http.StatusBadRequest)
		return
	} else if !isValidPhone(contact.Phone) {
		http.Error(w, "Number should be like: +7 (800) 555-35-55", http.StatusBadRequest)
	}

	contact.Id = id
	if err := h.service.UpdateContact(contact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, "Error while encoding contact", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (h *handler) deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteContact(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"contact": "deleted"})
	if err != nil {
		http.Error(w, "Error while encoding contact", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
