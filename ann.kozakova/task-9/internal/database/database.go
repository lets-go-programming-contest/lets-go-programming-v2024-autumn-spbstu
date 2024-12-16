package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/nutochk/task-9/internal/config"
)

type Database struct {
	DB *pgxpool.Pool
}

func NewDB(config *config.Config) (*Database, error) {
	conStr := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		config.DataBase.User,
		config.DataBase.Password,
		config.DataBase.Host,
		config.DataBase.Port,
		config.DataBase.Name,
	)

	pool, err := pgxpool.Connect(context.Background(), conStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %vn", err)
	}

	pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return &Database{pool}, nil
}