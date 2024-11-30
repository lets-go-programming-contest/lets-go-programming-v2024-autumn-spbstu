package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ErrEnv = errors.New("LoadConfig error")
)

const (
	envPath = ".env"
)

type AppCfg struct {
	Host string
	DBCfg
}

type DBCfg struct {
	UDBName     string
	UDBPass     string
	PgSQLHost   string
	DBPgSQLName string
	PortPgSQL   int
}

func LoadConfig() (AppCfg, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return AppCfg{}, fmt.Errorf("%w: %w", ErrEnv, err)
	}

	appConfigs := AppCfg{Host: os.Getenv("APP_PORT")}

	dbConfigs := DBCfg{
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
