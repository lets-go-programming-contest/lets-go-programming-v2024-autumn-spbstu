package db

import (
	"database/sql"
	"fmt"

	"github.com/artem6554/task-9/internal/config"
	_ "github.com/lib/pq"
)

func ConnectDB(data config.DbData) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", data.User, data.Password, data.Name) //TODO: Только здесь задаются данный для подключения к бд, надо вынести в конфиг
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
