package main

import (
	"log"
	"net/http"

	"github.com/IDevFrye/task-9/cmd/server/internal/config"
	"github.com/IDevFrye/task-9/cmd/server/internal/db"
	"github.com/IDevFrye/task-9/internal/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	configPath := config.GetConfigPath()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	database, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer database.Close()

	controllers.InitDB(database)
	router := mux.NewRouter()
	controllers.RegisterRoutes(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	loggedRouter := handlers.CORS(headers, methods, origins)(router)

	port := "8080"
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, loggedRouter))
}
