package main

import (
	"log"
	"net/http"

	"github.com/IDevFrye/task-9/internal/controllers"
	"github.com/IDevFrye/task-9/internal/db"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database, err := db.ConnectDB()
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
