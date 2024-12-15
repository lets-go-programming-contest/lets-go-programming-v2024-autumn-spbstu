package http

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"task-9-1/internal/database"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var phoneRegexp = regexp.MustCompile("^\\+?\\d[ ]?[-(]?\\d{3}[-)]?[ ]?\\d{3}[- ]?\\d{2}[- ]?\\d{2}$")

func CheckPhone(phone string) bool {
	return phoneRegexp.MatchString(phone)
}

var db *database.DataBase

func InitDataBase(other *database.DataBase) {
	db = other
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if contact.Name == "" || contact.Phone == "" || !CheckPhone(contact.Phone) {
		http.Error(w, "Wrong input or no input is required", http.StatusBadRequest)
		return
	}

	var existingID int
	db.DB.QueryRow(context.Background(), "SELECT id FROM contact WHERE phone=$1", contact.Phone).Scan(&existingID)
	if existingID != 0 {
		http.Error(w, "Contact already exists", http.StatusConflict)
		return
	}

	var id int
	err = db.DB.QueryRow(context.Background(), `INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id`, contact.Name, contact.Phone).Scan(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	contact.ID = id
	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var contact Contact
	err = db.DB.QueryRow(context.Background(), "SELECT * FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.Name, &contact.Phone)
	if contact.ID == 0 || err != nil {
		http.Error(w, "can't find a contact", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "error while encoding in get-method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.DB.Query(context.Background(), `SELECT id, name, phone FROM contacts`)
	if err != nil {
		http.Error(w, "error while getting contacts from database", http.StatusInternalServerError)
		return
	}

	defer rows.Close()
	var contacts []Contact
	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			http.Error(w, "error while scanning contacts from rows", http.StatusInternalServerError)
			return
		}
		contacts = append(contacts, contact)
	}
	err = json.NewEncoder(w).Encode(contacts)
	if err != nil {
		http.Error(w, "error while encoding contacts", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res, err := db.DB.Exec(context.Background(), `DELETE FROM contacts`)
	if err != nil {
		http.Error(w, "error while deleting contacts", http.StatusInternalServerError)
		return
	}

	con := res.RowsAffected()
	if con == 0 {
		http.Error(w, "no contacts deleted", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]int64{"deleted": con})
	if err != nil {
		http.Error(w, "error while encoding data", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID in update-method", http.StatusBadRequest)
		return
	}

	var contact Contact
	err = json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "error while decoding", http.StatusBadRequest)
		return
	}

	if contact.Name == "" || contact.Phone == "" || !CheckPhone(contact.Phone) {
		http.Error(w, "Wrong input or no input is required", http.StatusBadRequest)
		return
	}

	cont, err := db.DB.Exec(context.Background(), "UPDATE contacts SET name =$1, phone =$2 WHERE id = $3", contact.Name, contact.Phone, contactID)
	if err != nil {
		http.Error(w, "error while updating contact", http.StatusInternalServerError)
		return
	}

	rowsAffected := cont.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "error while updating contact", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "contact has been updated"})
	if err != nil {
		http.Error(w, "error while encoding in update-method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID in delete-method", http.StatusBadRequest)
		return
	}

	cont, err := db.DB.Exec(context.Background(), "DELETE FROM contacts WHERE id = $1", contactID)
	if err != nil {
		http.Error(w, "error while deleting contact", http.StatusInternalServerError)
		return
	}

	rowsAffected := cont.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "error while deleting contact", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "contact has been deleted"})
	if err != nil {
		http.Error(w, "error while encoding in delete-method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
