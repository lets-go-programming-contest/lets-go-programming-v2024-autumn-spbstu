package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DBstruct struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type DataBase struct {
	DB *sql.DB
}

func ConnectDB(cfg DBstruct) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDatabaseConnect, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%w: %w", ErrDatabasePing, err)
	}

	log.Println("Connect to the database successfully")
	return db, nil
}
