package db

import (
	"database/sql"
	"fmt"
	"task-9/internal/config"
	"task-9/internal/models"
	"task-9/internal/userErrors"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(cfg config.Config) (*sql.DB, error) {

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", userErrors.ErrInitDB, err)
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("%w: %w", userErrors.ErrInitDB, err)
	}

	//fmt.Println("Successfully connected to the database")
	return db, nil
}

func GetAllContacts() ([]models.Contact, error) {
	rows, err := db.Query("SELECT id, name, phone FROM contacts")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []models.Contact
	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}
	return contacts, nil
}

func GetContactByID(id int) (models.Contact, error) {
	var contact models.Contact
	err := db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id).Scan(&contact.ID, &contact.Name, &contact.Phone)
	if err != nil {
		return contact, err
	}
	return contact, nil
}

func CreateContact(contact *models.Contact) error {

	var existingContactID int
	errExistringContactID := db.QueryRow("SELECT id FROM contacts WHERE phone = $1", contact.Phone).Scan(&existingContactID)
	if errExistringContactID != nil || existingContactID != 0 {
		return errExistringContactID
	}

	err := db.QueryRow("INSERT INTO contacts (name, phone) VALUES ($1, $2) RETURNING id", contact.Name, contact.Phone).Scan(&contact.ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateContact(contact *models.Contact) error {
	_, err := db.Exec("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3", contact.Name, contact.Phone, contact.ID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteContact(id int) error {
	_, err := db.Exec("DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func CloseDB(db *sql.DB) error {

	err := db.Close()

	if err != nil {
		return err
	}

	return nil
}
