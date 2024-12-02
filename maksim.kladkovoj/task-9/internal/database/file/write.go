package database

import (
	"fmt"
	"log"
	"regexp"
)

type WriteDB struct {
	DataBase
}

var regexpPhone *regexp.Regexp

func NewDatabaseWriter(cfg DBstruct) (*WriteDB, error) {
	db, err := ConnectDB(cfg)
	if err != nil {
		return nil, err
	}

	return &WriteDB{
		DataBase: DataBase{
			DB: db,
		},
	}, nil
}

func isValidNumber(phone string) error {
	regexpPhone = regexp.MustCompile(`^(\+7|8)[-\s]?\(?\d{3}\)?[-\s]?\d{3}[-\s]?\d{2}[-\s]?\d{2}$`)

	if !regexpPhone.MatchString(phone) {
		return ErrPhoneInvalid
	}
	return nil
}

func (w *WriteDB) Post(contact Contact) error {
	if err := isValidNumber(contact.Phone); err != nil {
		return err
	}

	log.Println(contact)

	if err := w.DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect db: %w", err)
	}

	if _, err := w.DB.Query("INSERT INTO contacts (id, name, phone) values ($1, $2, $3)",
		contact.ID, contact.Name, contact.Phone); err != nil {
		return fmt.Errorf("%w: %w", ErrDatabaseQuery, err)
	}

	return nil
}

func (w *WriteDB) Put(contact Contact) error {
	if err := isValidNumber(contact.Phone); err != nil {
		return err
	}

	if _, err := w.DB.Query("UPDATE contacts SET name = $1, phone = $2 WHERE id = $3",
		contact.Name, contact.Phone, contact.ID); err != nil {
		return fmt.Errorf("%w: %w", ErrDatabaseQuery, err)
	}

	return nil
}

func (w *WriteDB) Delete(id int) error {
	if _, err := w.DB.Query("DELETE FROM contacts WHERE id = $1", id); err != nil {
		return fmt.Errorf("%w: %w", ErrDatabaseQuery, err)
	}

	return nil
}
