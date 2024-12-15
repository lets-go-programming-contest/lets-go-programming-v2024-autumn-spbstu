package database

import (
	"fmt"
	"task-9-1/internal/config"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
)

type DataBase struct {
	DB *pgxpool.Pool
}

func CreateDb(cfg config.DataBaseCfg) (*DataBase, error) {
	connection := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.HostDB,
		cfg.PortDB,
		cfg.NameDB)
	pool, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}

	return &DataBase{DB: pool}, nil
}
