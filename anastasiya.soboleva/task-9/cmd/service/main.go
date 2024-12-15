package main

import (
	"flag"
	"fmt"
	"github.com/nayzzerr/task-9/internal/config"
	"github.com/nayzzerr/task-9/internal/db"
	"github.com/nayzzerr/task-9/internal/routes"
	"log"
	"net/http"
)

func main() {
	configPath := flag.String("config", "config.json", "Path to the configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	if cfg.Server.Port == 0 {
		log.Fatalf("Server port is not specified in the configuration.")
	}

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer database.Close()

	router := routes.SetupRoutes(database)

	log.Printf("Server started on port %d", cfg.Server.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
