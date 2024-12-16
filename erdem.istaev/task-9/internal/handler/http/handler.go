package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"erdem.istaev/task-9/internal/service"
	"erdem.istaev/task-9/internal/structure"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service, r *mux.Router) *mux.Router {
	h := &Handler{services: services}

	r.HandleFunc("/upload", h.createContact).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/delete/{id}", h.deleteContact).Methods(http.MethodOptions, http.MethodDelete)
	r.HandleFunc("/get/{id}", h.getContactById).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/get", h.getAllContacts).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/put/{id}", h.updateContact).Methods(http.MethodOptions, http.MethodPut)

	return r
}

func (h *Handler) createContact(w http.ResponseWriter, r *http.Request) {
	var contact structure.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, fmt.Errorf("error while decode contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	id, err := h.services.CreateContact(contact)
	if err != nil {
		http.Error(w, fmt.Errorf("error while create contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	contact.Id = id
	if err = json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, fmt.Errorf("error while encode contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Errorf("error while parse id, invalid id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	contact, err := h.services.GetContactById(id)
	if err != nil {
		http.Error(w, fmt.Errorf("error while get contact by id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, fmt.Errorf("error while encode contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) getAllContacts(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.services.GetAllContacts()
	if err != nil {
		http.Error(w, fmt.Errorf("error while get all contacts: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err = json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, fmt.Errorf("error while encode all contacts: %w", err).Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Errorf("error while parse id, invalid id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err = h.services.DeleteContact(id); err != nil {
		http.Error(w, fmt.Errorf("error while delete contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) updateContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Errorf("error while parse id, invalid id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	var contact structure.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, fmt.Errorf("error while decode contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err := h.services.UpdateContact(id, contact); err != nil {
		http.Error(w, fmt.Errorf("error while update contact: %w", err).Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
