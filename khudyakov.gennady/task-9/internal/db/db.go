package db

import (
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/net/context"
)

func NewDB(cfg DBConfig, attempts int) (*pgxpool.Pool, error) {
	connect := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name)

	pool, err := pgxpool.Connect(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return pool, nil
}
