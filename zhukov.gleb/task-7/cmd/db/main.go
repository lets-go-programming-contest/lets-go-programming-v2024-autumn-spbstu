package main

import (
	"database/sql"
	"fmt"
	"log"

	dbPack "task-7/internal/db"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=username dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbService := dbPack.New(db)

	names, err := dbService.GetNames()

	for _, name := range names {
		fmt.Println(name)
	}
}
