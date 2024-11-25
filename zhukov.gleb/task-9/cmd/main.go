package main

import (
	"net/http"

	"task-9/internal/contact"
	"task-9/internal/db"
	"task-9/internal/handler"

	"github.com/gorilla/mux"
)

func main() {
	config, err := db.LoadConfig()
	if err != nil {
		panic("error reading data .env file: " + err.Error())
	}

	pgSQL, err := db.NewPgSQLController(
		config.UDBName,
		config.UDBPass,
		config.PgSQLHost,
		config.DBPgSQLName,
		config.PortPgSQL,
	)
	if err != nil {
		panic("error with pgSQL DB: " + err.Error())
	}

	contactRepo := contact.NewContactRepo(pgSQL)
	contactHandler := &handler.ContactHandler{
		ContactRepo: contactRepo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/contacts", contactHandler.GetAll).Methods("GET")
	r.HandleFunc("/contacts/{id}", contactHandler.GetByID).Methods("GET")
	r.HandleFunc("/contacts", contactHandler.GetByID).Methods("POST")

	addr := "127.0.0.1:8080"
	err = http.ListenAndServe(addr, r)
	if err != nil {
		panic("ListenAndServe fail" + err.Error())
	}
}
