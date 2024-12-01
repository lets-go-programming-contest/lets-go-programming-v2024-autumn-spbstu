package main

import (
	"database/sql"
	"fmt"
	"log"

	dbPack "github.com/kirill.romanchuk/task-7/internal/db"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		log.Fatalf("Error getting names: %v", err)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	defer db.Close()
}
