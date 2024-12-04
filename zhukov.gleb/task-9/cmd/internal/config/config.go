package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"task-9/internal/db"

	"github.com/joho/godotenv"
)

var (
	ErrEnv = errors.New("LoadConfig error")
)

// TODO env into configs dir
const (
	envPath = ".env"
)

type AppCfg struct {
	Port  string
	Host  string
	DBCfg db.Cfg
}

func LoadConfig() (AppCfg, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return AppCfg{}, fmt.Errorf("%w: %w", ErrEnv, err)
	}

	appConfigs := AppCfg{
		Port: os.Getenv("APP_PORT"),
		Host: os.Getenv("APP_HOST"),
	}

	dbConfigs := db.Cfg{
		UDBPass:     os.Getenv("USER_PGSQL_PASS"),
		DBPgSQLName: os.Getenv("DB_PGSQL_NAME"),
		UDBName:     os.Getenv("USER_PGSQL_NAME"),
		PgSQLHost:   os.Getenv("PGSQL_HOST"),
	}

	dbConfigs.PortPgSQL, err = strconv.Atoi(os.Getenv("PORT_PGSQL"))
	if err != nil {
		return AppCfg{}, fmt.Errorf("%w: %w", ErrEnv, err)
	}

	appConfigs.DBCfg = dbConfigs

	return appConfigs, nil
}
