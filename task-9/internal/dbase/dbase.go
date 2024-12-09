package dbase

import (
	config "contactManager/internal/config"
	errors "contactManager/internal/errorsExt"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDB(config config.Config) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, errors.ErrorLocate(err)
	}

	if err = db.Ping(); err != nil {
		return nil, errors.ErrorLocate(err)
	}

	fmt.Println("The database is connected")
	return db, nil
}

func CloseDB(db *sql.DB) error {

	err := db.Close()
	if err != nil {
		return errors.ErrorLocate(err)
	}
	return nil
}
