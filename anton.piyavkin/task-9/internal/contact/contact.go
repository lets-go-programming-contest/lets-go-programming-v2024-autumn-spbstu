package contact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
)

type Contact struct {
	ID    int
	Name  string
	Phone string
}

var db *pgxpool.Pool

func InitDB(tmp *pgxpool.Pool) {
	db = tmp
}

func isNum(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func isValidPhoneNumber(phone string) bool {
	if len(phone) < 1 || phone[0] != '+' {
		return false
	}

	parts := strings.Fields(phone[1:])
	if len(parts) != 3 {
		return false
	}

	countryCode := parts[0]
	if len(countryCode) < 1 || len(countryCode) > 3 || !isNum(countryCode) {
		return false
	}

	thirdPart := strings.Split(parts[2], "-")
	if len(thirdPart) != 3 || len(parts[1]) != 3 || len(thirdPart[0]) != 3 || !isNum(parts[1]) || !isNum(thirdPart[0]) {
		fmt.Println(4)
		return false
	}
	if len(thirdPart[1]) != 2 || !isNum(thirdPart[1]) || !isNum(thirdPart[2]) {
		return false
	}

	return true
}

func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var con Contact
	err := json.NewDecoder(r.Body).Decode(&con)
	if err != nil {
		http.Error(w, "contact not created", http.StatusBadRequest)
		return
	} else if con.Phone == "" || con.Name == "" || !isValidPhoneNumber(con.Phone) {
		http.Error(w, "incorrect name or phone number", http.StatusBadRequest)
		return
	}

	var existID int
	db.QueryRow(context.Background(), `SELECT id FROM contacts WHERE phone = $1`, con.Phone).Scan(&existID)
	if existID != 0 {
		http.Error(w, "existing number", http.StatusConflict)
		return
	}

	var createId int
	err = db.QueryRow(context.Background(), `INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id`, con.Name, con.Phone).Scan(&createId)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	con.ID = createId
	err = json.NewEncoder(w).Encode(con)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	conID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "incorrect ID", http.StatusBadRequest)
		return
	}

	var con Contact
	err = json.NewDecoder(r.Body).Decode(&con)
	if err != nil {
		http.Error(w, "invalid decode", http.StatusBadRequest)
		return
	}

	if con.Name == "" || con.Phone == "" || !isValidPhoneNumber(con.Phone) {
		http.Error(w, "incorrect name or phone number", http.StatusBadRequest)
		return
	}

	ct, err := db.Exec(context.Background(), `UPDATE contacts SET name = $1, phone = $2 WHERE id = $3`, con.Name, con.Phone, conID)
	if err != nil {
		http.Error(w, "failed update", http.StatusInternalServerError)
		return
	}

	changed := ct.RowsAffected()
	if changed == 0 {
		http.Error(w, "contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"contact": "updated"})
	if err != nil {
		http.Error(w, "failed encode", http.StatusInternalServerError)
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	conID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "incorrect ID", http.StatusBadRequest)
		return
	}

	ct, err := db.Exec(context.Background(), `DELETE FROM contacts WHERE id = $1`, conID)
	if err != nil {
		http.Error(w, "failed delete", http.StatusInternalServerError)
		return
	}

	changed := ct.RowsAffected()
	if changed == 0 {
		http.Error(w, "contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"contact": "deleted"})
	if err != nil {
		http.Error(w, "failed encode", http.StatusInternalServerError)
		return
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	conID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "incorrect ID", http.StatusBadRequest)
		return
	}

	var con Contact
	err = db.QueryRow(context.Background(), `SELECT * FROM contacts WHERE id = $1`, conID).Scan(&con.ID, &con.Name, &con.Phone)
	if con.ID == 0 || err != nil {
		http.Error(w, "contact not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(con)
	if err != nil {
		http.Error(w, "failed encode", http.StatusInternalServerError)
		return
	}
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query(context.Background(), `SELECT id, name, phone FROM contacts`)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var cons []Contact
	for rows.Next() {
		var con Contact
		err = rows.Scan(&con.ID, &con.Name, &con.Phone)
		if err != nil {
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}
		cons = append(cons, con)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(cons)
	if err != nil {
		http.Error(w, "failed encode", http.StatusInternalServerError)
		return
	}
}
