package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"task-9/internal/db"
	"task-9/internal/models"

	"github.com/gorilla/mux"
)

func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	contacts, err := db.GetAllContacts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	contact, err := db.GetContactByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var contact models.Contact

	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.ValidateContact(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateContact(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var contact models.Contact
	err = json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := models.ValidateContact(&contact); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	contact.ID = id
	err = db.UpdateContact(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(contact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	err = db.DeleteContact(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
