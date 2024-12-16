package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/EmptyInsid/task-9/internal/config"
)

type Database struct {
	pool *pgxpool.Pool
}

func NewDB(config *config.DatabaseConfig) (*Database, error) {
	db := &Database{pool: nil}
	if err := db.Init(buildConnectionString(config)); err != nil {
		return nil, err
	}

	// chekc connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.Pool().Ping(ctx); err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return db, nil
}

func (db *Database) Init(dsn string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var err error

	db.pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		return fmt.Errorf("%w", err)
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
