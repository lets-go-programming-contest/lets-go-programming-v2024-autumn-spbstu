package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nutochk/task-9/internal/database"
)

type Contact struct {
	ID        int `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
}

var db *database.Database

func InitDataBase(datab *database.Database){
	db = datab
}

var phoneRegexp = regexp.MustCompile("^\\+(\\d{1,3})\\s\\d{3}\\s\\d{3}-\\d{2}-\\d{2}$")

func isValidPhone(phone string) bool {
	return phoneRegexp.MatchString(phone)
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.DB.Query(context.Background(),"SELECT id, name, phone FROM contacts")
	if err != nil {
		http.Error(w, "Failed to fetch contacts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		contact := Contact{}
		rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		contacts = append(contacts, contact)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		http.Error(w, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var contact Contact
	db.DB.QueryRow(context.Background(),"SELECT * FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.Name, &contact.Phone)
	if contact.ID == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if contact.Name == "" || contact.Phone == "" {
		http.Error(w, "Name and phone are required", http.StatusBadRequest)
		return
	}

	if !isValidPhone(contact.Phone) {
		http.Error(w, "Invalid phone number format. Example: +7 123 456-78-90", http.StatusBadRequest)
		return
	}

	var existingContactID int
	db.DB.QueryRow(context.Background(),"SELECT id FROM contacts WHERE phone = $1", contact.Phone).Scan(&existingContactID)
	if existingContactID != 0 {
		http.Error(w, "Contact already exists", http.StatusConflict)
		return
	}

	var newID int
	err = db.DB.QueryRow(context.Background(),"INSERT INTO contacts (name, phone) " +
		"VALUES ($1, $2) RETURNING id", contact.Name, contact.Phone).
		Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create contact", http.StatusInternalServerError)
		return
	}

	contact.ID = newID
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var contact Contact
	err = json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if contact.Name == "" || contact.Phone == "" {
		http.Error(w, "Name and phone are required", http.StatusBadRequest)
		return
	}

	if !isValidPhone(contact.Phone) {
		http.Error(w, "Invalid phone number format. Example: +7 (123) 456-78-90", http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(context.Background(), "UPDATE contacts SET name = $1, phone = $2 " +
		"WHERE id = $3", contact.Name, contact.Phone, contactID)

	if err != nil {
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Contact updated successfully"})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
	}

	res, err := db.DB.Exec(context.Background(),"DELETE FROM contacts WHERE id = $1", contactID)
	if err != nil {
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}

	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"message": "Contact deleted successfully"})
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}