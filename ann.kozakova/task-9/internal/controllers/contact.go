package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"

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

var phoneRegexp = regexp.MustCompile("^\\+(\\d{1,3})\\s\\(\\d{3}\\)\\s\\d{3}-\\d{2}-\\d{2}$")

func isValidPhone(phone string) bool {
	return phoneRegexp.MatchString(phone)
}

func GetContacts(res http.ResponseWriter, req *http.Request) {
	rows, err := db.DB.Query(context.Background(),"SELECT id, name, phone FROM contacts")
	if err != nil {
		http.Error(res, "Failed to fetch contacts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		contact := Contact{}
		rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		contacts = append(contacts, contact)
	}

	err = json.NewEncoder(res).Encode(contacts)
	if err != nil {
		http.Error(res, "Failed to encode contact", http.StatusInternalServerError)
		return
	}
}