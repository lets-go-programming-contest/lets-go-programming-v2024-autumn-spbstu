package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/nutochk/task-9/internal/config"
	"github.com/nutochk/task-9/internal/controllers"
	"github.com/nutochk/task-9/internal/database"
)

func main() {
	config, err := config.ReadConfig(config.GetPath())
	if err != nil {
		panic(err)
	}
	db, err := database.NewDB(config)
	defer db.DB.Close()
	controllers.InitDataBase(db)
	router := mux.NewRouter()
	controllers.CreateRoutes(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	loggedRouter := handlers.CORS(headers, methods, origins)(router)

	port := config.Server.Port
	log.Fatal(http.ListenAndServe(":"+port, loggedRouter))
}