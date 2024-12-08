package main

import (
	"fmt"
	"log"

	db "github.com/Koshsky/task-9/internal/database"
)

func main() {
	// Параметры подключения
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "contact_db"

	cm, err := db.NewContactManager(host, port, user, password, dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer cm.Close()

	// Пример использования CRUD операций
	contactID, err := cm.CreateContact("John Doe", "123-456-7890")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Создан контакт с ID: %d\n", contactID)

	contact, err := cm.GetContact(contactID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Получен контакт: %+v\n", contact)

	err = cm.UpdateContact(contactID, "Jane Doe", "098-765-4321")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Контакт обновлен")

	err = cm.DeleteContact(contactID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Контакт удален")
}
