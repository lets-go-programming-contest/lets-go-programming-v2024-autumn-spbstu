package main

import (
	"log"
	"net/http"
	"task-9-1/internal/config"
	"task-9-1/internal/database"
	http2 "task-9-1/internal/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	path := config.GetPathOfFile()
	cfg, err := config.ParseConfig(path)
	if err != nil {
		panic(err)
	}

	db, err := database.CreateDb(cfg.DataBase)
	if err != nil {
		panic(err)
	}
	defer db.DB.Close()

	http2.InitDataBase(db)
	router := mux.NewRouter()
	http2.CreateRoutes(router)

	logs := handlers.CORS()(router)

	port := cfg.Server.Port
	log.Fatal(http.ListenAndServe(":"+port, logs))
}
