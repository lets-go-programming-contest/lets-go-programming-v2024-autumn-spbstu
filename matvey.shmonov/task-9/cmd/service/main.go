package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/Koshsky/task-9/internal/api"
	"github.com/Koshsky/task-9/internal/config"
	db "github.com/Koshsky/task-9/internal/database"
	"github.com/gorilla/mux"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "./config/config.json", "config file path")
	flag.Parse()

	config, _ := config.LoadConfig(path)
	cm, err := db.NewContactManager(
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cm.Close()
	r := mux.NewRouter()
	api.RegisterRoutes(r, cm)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server : %v", err)
	}
}
