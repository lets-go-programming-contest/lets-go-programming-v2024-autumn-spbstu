package http

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"task-9-1/database"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

var phoneRegexp = regexp.MustCompile("\\+?\\d[ ]?[-(]?\\d{3}[-)]?[ ]?\\d{3}[- ]?\\d{2}[- ]?\\d{2}")

func CheckPhone(phone string) bool {
	return phoneRegexp.MatchString(phone)
}

var db *database.DataBase

func InitDataBase(other *database.DataBase) {
	db = other
}

func CreateContact(w http.ResponseWriter, r http.Request) {
	w.Header().Set("Counter-Type", "application/json")
	var contact Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if CheckPhone(contact.Phone) || contact.Name == "" || contact.Phone == "" {
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
	w.Header().Set("Counter-Type", "application/json")

	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var contact Contact
	err = db.DB.QueryRow(context.Background(), "SELECT * FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.Name, &contact.Phone)
	if contact.ID == 0 || err != nil {
		http.Error(w, "error while getting a contact", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(contact)
	if err != nil {
		http.Error(w, "error while encoding in get-method", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
