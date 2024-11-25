package controllers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB

func InitDB(database *sql.DB) {
	db = database
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, phone, created_at, updated_at FROM contacts")
	if err != nil {
		http.Error(w, "Failed to fetch contacts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.CreatedAt, &contact.UpdatedAt)
		if err != nil {
			http.Error(w, "Failed to parse contacts", http.StatusInternalServerError)
			return
		}
		contacts = append(contacts, contact)
	}

	json.NewEncoder(w).Encode(contacts)
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
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

	var newID int
	err = db.QueryRow(`INSERT INTO contacts (name, phone) 
                        VALUES ($1, $2) RETURNING id`, contact.Name, contact.Phone).
		Scan(&newID)
	if err != nil {
		http.Error(w, "Failed to create contact", http.StatusInternalServerError)
		return
	}

	contact.ID = newID
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(contact)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
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

	// Проверяем, что имя и телефон не пустые
	if contact.Name == "" || contact.Phone == "" {
		http.Error(w, "Name and phone are required", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`UPDATE contacts SET name = $1, phone = $2, updated_at = CURRENT_TIMESTAMP 
                            WHERE id = $3`, contact.Name, contact.Phone, contactID)
	if err != nil {
		http.Error(w, "Failed to update contact", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact updated successfully"})
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	result, err := db.Exec(`DELETE FROM contacts WHERE id = $1`, contactID)
	if err != nil {
		http.Error(w, "Failed to delete contact", http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Contact deleted successfully"})
}

func GetTagsForContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	rows, err := db.Query(`SELECT t.id, t.name 
                           FROM tags t
                           JOIN contact_tags ct ON t.id = ct.tag_id
                           WHERE ct.contact_id = $1`, contactID)
	if err != nil {
		http.Error(w, "Failed to fetch tags", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			http.Error(w, "Failed to parse tags", http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}
	if len(tags) == 0 {
		tags = []Tag{}
	}

	json.NewEncoder(w).Encode(tags)
}

func AddTagsToContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var tagIDs []int
	body, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(body, &tagIDs)
	if err != nil || len(tagIDs) == 0 {
		http.Error(w, "Invalid tag data", http.StatusBadRequest)
		return
	}

	for _, tagID := range tagIDs {
		_, err := db.Exec(`INSERT INTO contact_tags (contact_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, contactID, tagID)
		if err != nil {
			http.Error(w, "Failed to associate tag", http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func RemoveTagFromContact(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contactID, err := strconv.Atoi(vars["id"])
	tagID, errTag := strconv.Atoi(vars["tag_id"])

	if err != nil || errTag != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(`DELETE FROM contact_tags WHERE contact_id = $1 AND tag_id = $2`, contactID, tagID)
	if err != nil {
		http.Error(w, "Failed to remove tag", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		http.Error(w, "Failed to fetch tags", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(&tag.ID, &tag.Name)
		if err != nil {
			http.Error(w, "Failed to parse tags", http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}

func GetContactsByTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tagIdStr := vars["tag_id"]
	tagId, err := strconv.Atoi(tagIdStr)
	if err != nil {
		http.Error(w, "Invalid tag ID", http.StatusBadRequest)
		return
	}

	rows, err := db.Query(`SELECT c.id, c.name, c.phone, c.created_at, c.updated_at 
                            FROM contacts c
                            JOIN contact_tags ct ON c.id = ct.contact_id
                            JOIN tags t ON ct.tag_id = t.id
                            WHERE t.id = $1`, tagId)
	if err != nil {
		http.Error(w, "Failed to fetch contacts by tag", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone, &contact.CreatedAt, &contact.UpdatedAt)
		if err != nil {
			http.Error(w, "Failed to parse contacts", http.StatusInternalServerError)
			return
		}
		contacts = append(contacts, contact)
	}

	if len(contacts) == 0 {
		contacts = []Contact{}
	}

	json.NewEncoder(w).Encode(contacts)
}
