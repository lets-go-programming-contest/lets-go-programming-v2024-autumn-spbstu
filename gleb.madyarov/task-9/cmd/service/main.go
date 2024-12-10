package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/Madyarov-Gleb/task-9/internal/config"
	"github.com/Madyarov-Gleb/task-9/internal/db"
	"github.com/Madyarov-Gleb/task-9/internal/handlers"
)

func main() {
	log.Println("Start server...")

	config, err := config.ReadConfig(config.GetPath())
	if err != nil {
		panic(err)
	}

	dbConn := db.ConnectionDatabase(config)
	defer dbConn.Close()

	db.InitDatabase(dbConn)

	router := mux.NewRouter()
	handlers.CreateRoutes(router, dbConn)

	portServer := config.PortServer
	log.Println("The server is running on port " + portServer)
	log.Fatal(http.ListenAndServe(":"+portServer, router))
}
