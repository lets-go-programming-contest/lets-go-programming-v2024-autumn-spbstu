package db

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	envPath = ".env"
)

var (
	ErrOpenDB   = errors.New("error open")
	ErrPingDB   = errors.New("error ping")
	ErrInsertDB = errors.New("error insert")
)

type DBcfg struct {
	UDBName     string
	UDBPass     string
	PgSQLHost   string
	DBPgSQLName string
	PortPgSQL   int
}

func LoadConfig() (DBcfg, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return DBcfg{}, err
	}

	dbConfigs := DBcfg{
		UDBPass:     os.Getenv("USER_PGSQL_PASS"),
		DBPgSQLName: os.Getenv("DB_PGSQL_NAME"),
		UDBName:     os.Getenv("USER_PGSQL_NAME"),
		PgSQLHost:   os.Getenv("PGSQL_HOST"),
	}

	dbConfigs.PortPgSQL, err = strconv.Atoi(os.Getenv("PORT_PGSQL"))
	if err != nil {
		return DBcfg{}, err
	}

	return dbConfigs, nil
}
