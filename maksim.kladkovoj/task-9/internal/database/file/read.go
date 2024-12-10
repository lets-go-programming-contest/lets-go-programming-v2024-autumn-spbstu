package database

import (
	"fmt"
)

type Contact struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ReadDB struct {
	DataBase
}

func NewDatabaseReader(cfg DBstruct) (*ReadDB, error) {
	db, err := ConnectDB(cfg)
	if err != nil {
		return nil, err
	}

	return &ReadDB{
		DataBase: DataBase{
			DB: db,
		},
	}, nil
}

func (r *ReadDB) Get(query string) ([]Contact, error) {
	data, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDatabaseQuery, err)
	}
	defer data.Close()

	var contacts []Contact

	for data.Next() {
		var contact Contact
		if err := data.Scan(&contact.ID, &contact.Name, &contact.Phone); err != nil {
			return nil, fmt.Errorf("%w: %w", ErrScanContact, err)
		}

		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func (r *ReadDB) Exists(query string) error {
	var exists bool

	err := r.DB.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrDatabaseQuery, err)
	}

	if !exists {
		return ErrContactNotExists
	}

	return nil
}
