package main

import (
	"log"
	"net/http"
	"task-9-1/config"
	"task-9-1/database"
	myhttp "task-9-1/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	cfg, err := config.ParseConfig("config/config.yaml") // добавить флаг
	if err != nil {
		panic(err)
	}

	db, err := database.CreateDb(cfg.DataBase)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	myhttp.InitDataBase(db)
	router := mux.NewRouter()
	myhttp.CreateRoutes(router)

	headers := handlers.AllowedHeaders([]string{"Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})

	origins := handlers.AllowedOrigins([]string{"http://localhost:8080"}) //fix
	logs := handlers.CORS(headers, methods, origins)(router)

	port := cfg.Server.Port
	log.Fatal(http.ListenAndServe(":"+port, logs))
}
