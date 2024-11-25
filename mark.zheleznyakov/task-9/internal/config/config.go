package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func loadDotenv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

func LoadPgConfig() (string, error) {
	err := loadDotenv()
	if err != nil {
		return "", err
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	port := os.Getenv("POSTGRES_PORT")
	tz := os.Getenv("POSTGRES_TZ")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		host,
		user,
		password,
		db,
		port,
		tz,
	), nil
}
