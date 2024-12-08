package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Contact struct {
	ID    int
	Name  string
	Phone string
}

type ContactManager struct {
	db *sql.DB
}

func NewContactManager(host, port, user, password, dbname string) (*ContactManager, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return &ContactManager{db: db}, nil
}

func (cm *ContactManager) CreateContact(name, phone string) (int, error) {
	var id int
	err := cm.db.QueryRow("INSERT INTO contacts(name, phone) VALUES($1, $2) RETURNING id", name, phone).Scan(&id)
	return id, err
}

func (cm *ContactManager) GetContact(id int) (*Contact, error) {
	contact := &Contact{}
	err := cm.db.QueryRow("SELECT id, name, phone FROM contacts WHERE id = $1", id).Scan(&contact.ID, &contact.Name, &contact.Phone)
	if err != nil {
		return nil, err
	}
	return contact, nil
}

func (cm *ContactManager) GetContacts() ([]Contact, error) {
	var contacts []Contact                                      
	rows, err := cm.db.Query("SELECT id, name, phone FROM contacts") 
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contact Contact
		if err := rows.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
			return nil, err
		}
		contacts = append(contacts, contact) 
	}

	if err := rows.Err(); err != nil {
		return nil, err 
	}

	return contacts, nil
}

func (cm *ContactManager) UpdateContact(id int, name, phone string) error {
	_, err := cm.db.Exec("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3", name, phone, id)
	return err
}

func (cm *ContactManager) DeleteContact(id int) error {
	_, err := cm.db.Exec("DELETE FROM contacts WHERE id = $1", id)
	return err
}

func (cm *ContactManager) Close() error {
	return cm.db.Close()
}
