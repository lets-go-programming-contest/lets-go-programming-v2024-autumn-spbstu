package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/nayzzerr/task-9/internal/models"
	"github.com/nayzzerr/task-9/internal/services"
	"github.com/nayzzerr/task-9/internal/validation"
)

type ContactHandler struct {
	service *services.ContactService
}

func NewContactHandler(service *services.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

func (h *ContactHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	contacts, err := h.service.GetAllContacts()
	if err != nil {
		http.Error(w, "Failed to retrieve contacts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		http.Error(w, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
}

func (h *ContactHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	contact, err := h.service.GetContactByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve contact", http.StatusInternalServerError)
		return
	}
	if contact == nil {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
}

func (h *ContactHandler) Create(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validation.ValidateContactData(contact.Name, contact.Phone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.CreateContact(&contact); err != nil {
		http.Error(w, "Failed to create contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
}

func (h *ContactHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var contact models.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	contact.ID = id

	if err := validation.ValidateContactData(contact.Name, contact.Phone); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateContact(&contact); err != nil {
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *ContactHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if err := h.service.DeleteContact(id); err != nil {
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
