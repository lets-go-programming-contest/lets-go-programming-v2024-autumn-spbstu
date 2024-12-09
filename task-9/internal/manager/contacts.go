package manager

import (
	// "contactManager/internal/dbase"
	"database/sql"
	// "fmt"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

func isValidPhoneNumber(phone string) bool {

	re := regexp.MustCompile(`^\+\d{1,2} \(\d{3}\) \d{3}-\d{2}-\d{2}$`)
	return re.MatchString(phone)
}

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var db *sql.DB

func InitDataBase(database *sql.DB) {
	db = database
}

func GetContacts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("SELECT id, name, phone FROM contacts")
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
	db.QueryRow("SELECT * FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.Name, &contact.Phone)
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

	if !isValidPhoneNumber(contact.Phone) {
		http.Error(w, "Invalid phone number format. Example: +7 123 456-78-90", http.StatusBadRequest)
		return
	}

	var existingContactID int
	db.QueryRow("SELECT id FROM contacts WHERE phone = $1", contact.Phone).Scan(&existingContactID)
	if existingContactID != 0 {
		http.Error(w, "Contact already exists", http.StatusConflict)
		return
	}

	var newID int
	err = db.QueryRow("INSERT INTO contacts (id, name, phone) "+
		"VALUES ($1, $2, $3) RETURNING id", contact.ID, contact.Name, contact.Phone).
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

	if !isValidPhoneNumber(contact.Phone) {
		http.Error(w, "Invalid phone number format. Example: +7 (123) 456-78-90", http.StatusBadRequest)
		return
	}

	res, err := db.Exec("UPDATE contacts SET name = $1, phone = $2 "+
		"WHERE id = $3", contact.Name, contact.Phone, contactID)

	if err != nil {
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "wrong query, poo-poo-poo", http.StatusNotFound)
		return
	}
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

	res, err := db.Exec("DELETE FROM contacts WHERE id = $1", contactID)
	if err != nil {
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "wrong query, poo-poo-poo", http.StatusNotFound)
		return
	}
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
