package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
)

//go:embed config.json
var configFile string

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"database"`
}

func main() {

	var config Config
	err := json.Unmarshal([]byte(configFile), &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Database config:\n")
	fmt.Printf("Host: %s\n", config.Database.Host)
	fmt.Printf("Port: %d\n", config.Database.Port)
	fmt.Printf("User: %s\n", config.Database.User)
	fmt.Printf("Password: %s\n", config.Database.Password)
}
