package contacts

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/artem6554/task-9/internal/db"
)

type Contact struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Number string `json:"number"`
}

func GetContact(name string) ([]byte, error) {
	var сontacts []Contact
	db, err := db.ConnectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	queryString := fmt.Sprintf("select name, number from numbers where name = '%v'", name)
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var contact Contact
		rows.Scan(&contact.Name, &contact.Number)
		сontacts = append(сontacts, contact)
	}
	result, err := json.Marshal(сontacts)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func AddContact(name string, number string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	queryString := fmt.Sprintf("insert into numbers (name, number) values ('%v', '%v')", name, number)
	_, err = db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

func EditNumber(name string, number string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	queryString := fmt.Sprintf("UPDATE table_name SET number = '%v' WHERE name = '%v'", number, name)
	_, err = db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

func DeleteContact(name string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	queryString := fmt.Sprintf("DELETE FROM numbers WHERE name = '%v'", name)
	_, err = db.Query(queryString)
	if err != nil {
		return err
	}
	return nil
}

func Exists(name string) error {
	db, err := db.ConnectDB()
	if err != nil {
		return err
	}
	defer db.Close()
	queryString := fmt.Sprintf("select id, name, number from numbers where name = '%v'", name)
	rows, err := db.Query(queryString)
	if err != nil {
		return err
	}
	if rows.Next() {
		return errors.New("file already exists")
	}
	return nil

}
