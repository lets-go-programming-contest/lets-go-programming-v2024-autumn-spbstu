package contacts

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/hahapathetic/task-9/internal/commonErrors"
	"github.com/hahapathetic/task-9/internal/config"
	"github.com/hahapathetic/task-9/internal/database"
)

var dbConfig config.DbData

func init() {
	dbConfig = config.ReadDbConfig()
}

type Contact struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

func connectDB() (*sql.DB, error) {
	return database.ConnectDB(dbConfig)
}

func GetAllContacts() ([]byte, error) {
	db, err := connectDB()
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	queryString := "SELECT id, name, number FROM numbers"
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, fmt.Errorf("query database: %w", err)
	}
	defer rows.Close()

	var contacts []Contact
	for rows.Next() {
		var contact Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Number); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		contacts = append(contacts, contact)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	result, err := json.Marshal(contacts)
	if err != nil {
		return nil, fmt.Errorf("marshal contacts: %w", err)
	}
	return result, nil
}

func GetContactByID(id string) (Contact, error) {
	db, err := connectDB()
	if err != nil {
		return Contact{}, fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	queryString := fmt.Sprintf("SELECT id, name, number FROM numbers WHERE id = '%v'", id)
	row := db.QueryRow(queryString)

	var contact Contact
	if err := row.Scan(&contact.ID, &contact.Name, &contact.Number); err != nil {
		if err == sql.ErrNoRows {
			return Contact{}, commonErrors.ErrContactNotExists // Возвращаем ошибку, если контакт не найден
		}
		return Contact{}, fmt.Errorf("scan row: %w", err)
	}

	return contact, nil
}

func AddContact(name, number string) error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	if err := ExistsByNumber(number); err != nil {
		return fmt.Errorf("contact with this number already exists: %w", err)
	}

	queryString := "INSERT INTO numbers (name, number) VALUES ($1, $2)"
	_, err = db.Exec(queryString, name, number)
	if err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}

func EditNumberByID(id, name, number string) error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	queryString := fmt.Sprintf("UPDATE numbers SET name = '%v', number = '%v' WHERE id = '%v'", name, number, id)
	_, err = db.Exec(queryString)
	if err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}

func DeleteContactByID(id string) error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	queryString := fmt.Sprintf("DELETE FROM numbers WHERE id = '%v'", id)
	_, err = db.Exec(queryString)
	if err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}

func ExistsByNumber(number string) error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	queryString := fmt.Sprintf("SELECT 1 FROM numbers WHERE number = '%v' LIMIT 1", number)

	rows, err := db.Query(queryString)
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		return commonErrors.ErrContactAlreadyExists
	}

	return nil
}

func ExistsByID(id string) error {
	db, err := connectDB()
	if err != nil {
		return fmt.Errorf("connect to db: %w", err)
	}
	defer db.Close()

	queryString := fmt.Sprintf("SELECT 1 FROM numbers WHERE id = '%v' LIMIT 1", id)

	rows, err := db.Query(queryString)
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}

	return commonErrors.ErrContactNotExists
}

func IsCorrectNumber(number string) error {
	numberRegExp := regexp.MustCompile("^\\+(\\d{1,3})\\s\\d{3}\\s\\d{3}-\\d{2}-\\d{2}$")
	if numberRegExp.MatchString(number) {
		return nil
	}
	return commonErrors.ErrIncorrectNumber
}
