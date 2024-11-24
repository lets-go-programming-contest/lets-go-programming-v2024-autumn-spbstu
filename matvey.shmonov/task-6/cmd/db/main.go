package main

import (
	"database/sql"
	"fmt"
	"log"

	dbPack "github.com/Koshsky/task-6/internal/db"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbService := dbPack.New(db)

	if names, err := dbService.GetNames(); err == nil {
		for _, name := range names {
			fmt.Println(name)
		}
	}
}
