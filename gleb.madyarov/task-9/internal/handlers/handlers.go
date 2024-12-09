package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/Madyarov-Gleb/task-9/internal/models"

	"github.com/gorilla/mux"
)

func ValidatePhone(phone string) error {
	regex := regexp.MustCompile(`^\+7\d{10}$`)

	if !regex.MatchString(phone) {
		return ErrPhoneFormat
	}

	return nil
}

func GetContacts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, phone FROM contacts")
		if err != nil {
			log.Printf("error: %v: %v", ErrGet, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var contacts []models.Contact
		for rows.Next() {
			var contact models.Contact
			if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
				log.Printf("error: %v: %v", ErrGet, err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			contacts = append(contacts, contact)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contacts)
	}
}

func GetContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var contact models.Contact
		err := db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id).Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			log.Printf("error: %v: %v", ErrGet, err)
			http.Error(w, ErrGet.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contact)
	}
}

func CreateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var contact models.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			log.Printf("error: %v: %v", ErrCreate, err)
			http.Error(w, ErrDecode.Error(), http.StatusBadRequest)
			return
		}

		if err := ValidatePhone(contact.Phone); err != nil {
			log.Printf("Phone validation error: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err := db.QueryRow("INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id", contact.Name, contact.Phone).Scan(&contact.ID)
		if err != nil {
			log.Printf("error: %v: %v", ErrCreate, err)
			http.Error(w, ErrCreate.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contact)
	}
}

func UpdateContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		var contact models.Contact
		if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
			log.Printf("error: %v: %v", ErrUpdate, err)
			http.Error(w, ErrDecode.Error(), http.StatusBadRequest)
			return
		}

		if err := ValidatePhone(contact.Phone); err != nil {
			log.Printf("Phone validation error: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err := db.Exec("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3", contact.Name, contact.Phone, id)
		if err != nil {
			log.Printf("error: %v: %v", ErrUpdate, err)
			http.Error(w, ErrUpdate.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func DeleteContact(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := db.Exec("DELETE FROM contacts WHERE id = $1", id)
		if err != nil {
			log.Printf("error: %v: %v", ErrDelete, err)
			http.Error(w, ErrDelete.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
