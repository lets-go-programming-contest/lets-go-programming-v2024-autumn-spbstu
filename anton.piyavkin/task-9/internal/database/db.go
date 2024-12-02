package database

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"

	"github.com/Piyavva/task-9/internal/config"
)

func NewDB(cfg config.ConfigDB, attempts int) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.UserDB,
		cfg.PasswordDB,
		cfg.HostDB,
		cfg.PortDB,
		cfg.NameDB)

	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
