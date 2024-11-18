package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	dbPack "task-6/internal/db"
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
	if err != nil {
		fmt.Println(err)

		return
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
