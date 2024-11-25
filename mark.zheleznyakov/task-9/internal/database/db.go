package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mrqiz/task-9/internal/config"
)

var DB *gorm.DB

func Connect() error {
	var err error
	var connStr string
	connStr, err = config.LoadPgConfig()

	if err != nil {
		return fmt.Errorf("unable to read cfg: %d", err)
	}

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("unable to connect: %d", err)
	}

	return nil
}
