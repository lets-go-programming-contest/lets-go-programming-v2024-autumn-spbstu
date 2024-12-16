package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Madyarov-Gleb/task-9/internal/config"
)

func InitDatabase(db *sql.DB) {
	createContactsTable := `
	CREATE TABLE IF NOT EXISTS contacts (
		id SERIAL PRIMARY KEY,       
		name VARCHAR(100) NOT NULL,  
		phone VARCHAR(20) NOT NULL UNIQUE 
	);`

	_, err := db.Exec(createContactsTable)
	if err != nil {
		log.Fatalf("Error creating the 'contacts' table: %v", err)
	}

	log.Println("The 'contacts' table has been successfully created or already exists.")
}

func ConnectionDatabase(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.PortDb,
		cfg.DbUser,
		cfg.DbName,
		cfg.Password,
	)

	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	log.Println("Connection successful...")

	return dbConn
}
