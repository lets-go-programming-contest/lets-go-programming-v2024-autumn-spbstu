package database

import (
	"database/sql"
	"fmt"

	"github.com/hahapathetic/task-9/internal/config"
	_ "github.com/lib/pq"
)

func ConnectDB(data config.DbData) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", data.User, data.Password, data.Name)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
