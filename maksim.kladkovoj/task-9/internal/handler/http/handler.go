package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	database "github.com/Mmmakskl/task-9/internal/database/file"
	"github.com/gorilla/mux"
)

type service interface {
	Post(int, string, string) error
	Put(int, string, string) error
	Delete(int) error
	GetAll() ([]database.Contact, error)
	GetID(int) ([]database.Contact, error)
}

type handlers struct {
	service service
}

func NewHandler(s service, r *mux.Router) *mux.Router {
	h := handlers{
		service: s,
	}

	r.HandleFunc("/get/contacts", h.getAll).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/get/contacts/{id}", h.getID).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/post/contacts", h.post).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/put/contacts/{id}", h.put).Methods(http.MethodOptions, http.MethodPut)
	r.HandleFunc("/delete/contacts/{id}", h.delete).Methods(http.MethodOptions, http.MethodDelete)

	return r
}

func (h *handlers) getID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Errorf("invalid id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	contact, err := h.service.GetID(id)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrContactNotExists):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, fmt.Errorf("error while get contact id %d: %w", id, err).Error(), http.StatusNotFound)
		}
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")

	if err := json.NewEncoder(w).Encode(contact); err != nil {
		http.Error(w, fmt.Errorf("encode contacts failed: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handlers) getAll(w http.ResponseWriter, r *http.Request) {
	contacts, err := h.service.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, database.ErrContactNotExists):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, fmt.Errorf("error while get contacts: %w", err).Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/x-yaml")

	if err := json.NewEncoder(w).Encode(contacts); err != nil {
		http.Error(w, fmt.Errorf("encode contacts failed: %w", err).Error(), http.StatusInternalServerError)
		return
	}
}

func (h *handlers) post(w http.ResponseWriter, r *http.Request) {
	var contact database.Contact

	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, fmt.Errorf("invalid input YAML data: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Post(contact.ID, contact.Name, contact.Phone); err != nil {
		http.Error(w, fmt.Errorf("failed contact created: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *handlers) put(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Errorf("invalid id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	var contact database.Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		http.Error(w, fmt.Errorf("invalid input YAML data: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Put(contact.ID, contact.Name, contact.Phone); err != nil {
		switch {
		case errors.Is(err, database.ErrContactNotExists):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, fmt.Errorf("error while update contact id %d: %w", id, err).Error(), http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *handlers) delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, fmt.Errorf("invalid id: %w", err).Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id); err != nil {
		switch {
		case errors.Is(err, database.ErrContactNotExists):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, fmt.Errorf("error while delete contacct: %w", err).Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
