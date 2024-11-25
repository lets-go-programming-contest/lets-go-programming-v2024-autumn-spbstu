package main

import (
	"database/sql"
	"fmt"
	"log"

	dbPack "github.com/hahapathetic/task-7/internal/db"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Ошибка подключения к базе данных: %v", err)

		return
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		log.Printf("Ошибка получения имен: %v", err)

		return
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
