package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/mrqiz/task-9/internal/config"
)

var DB *gorm.DB

func Connect() {
	var err error
	var connStr string
	connStr, err = config.LoadPgConfig()

	if err != nil {
		log.Fatalf("unable to read cfg: %d", err)
	}

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("unable to connect: %d", err)
	}
}
