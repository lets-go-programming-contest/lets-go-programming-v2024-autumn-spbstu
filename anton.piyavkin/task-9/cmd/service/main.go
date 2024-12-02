package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/Piyavva/task-9/internal/config"
	"github.com/Piyavva/task-9/internal/contact"
	"github.com/Piyavva/task-9/internal/database"
	"github.com/Piyavva/task-9/internal/flag"
	"github.com/Piyavva/task-9/internal/route"
)

func main() {
	cfg := config.Load(flag.NameFile)
	db, err := database.NewDB(cfg.DataBase, 10)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	contact.InitDB(db)
	router := mux.NewRouter()
	route.Create(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	logs := handlers.CORS(headers, methods, origins)(router)

	port := cfg.Server.Port
	log.Fatal(http.ListenAndServe(":"+port, logs))
}
