package database

import (
	"context"
	"fmt"
	"time"

	"github.com/EmptyInsid/task-9/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDb(config *config.DatabaseConfig) (*Database, error) {
	db := &Database{}
	db.Init(buildConnectionString(config))

	// chekc connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Pool().Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) Init(dsn string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error
	db.pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) Close() {
	if db.pool != nil {
		db.pool.Close()
	}
}

func (db *Database) Pool() *pgxpool.Pool {
	return db.pool
}

func buildConnectionString(config *config.DatabaseConfig) string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}
