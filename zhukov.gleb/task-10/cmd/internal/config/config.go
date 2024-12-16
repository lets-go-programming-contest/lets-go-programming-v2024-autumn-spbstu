package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"task-10/internal/db"

	"github.com/joho/godotenv"
)

var (
	ErrEnv = errors.New("LoadConfig error")
)

const (
	envPath = ".env"
)

type AppCfg struct {
	RESTPort string
	GRPCPort string
	Host     string
	DBCfg    db.Cfg
}

func LoadConfig() (AppCfg, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return AppCfg{}, fmt.Errorf("%w: %w", ErrEnv, err)
	}

	appConfigs := AppCfg{
		RESTPort: os.Getenv("APP_REST_PORT"),
		GRPCPort: os.Getenv("APP_GRPC_PORT"),
		Host:     os.Getenv("APP_HOST"),
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
