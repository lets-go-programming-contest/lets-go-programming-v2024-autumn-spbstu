package main

import (
	"database/sql"
	"fmt"
	"log"

	dbPack "github.com/EmptyInsid/task-7/internal/db"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error while open db: %w", err)
	}

	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()
	if err != nil {
		return fmt.Errorf("error while GetNames db: %w", err)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
