package main

import (
	"fmt"
	"log"

	"github.com/Koshsky/task-9/internal/config"
	db "github.com/Koshsky/task-9/internal/database"
)

func main() {
	config, _ := config.LoadConfig("./config/config.json")
	cm, err := db.NewContactManager(
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
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
