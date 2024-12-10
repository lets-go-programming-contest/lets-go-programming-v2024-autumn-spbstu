package handlers

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router, dbConn *sql.DB) {
	router.HandleFunc("/contacts", GetContacts(dbConn)).Methods("GET")
	router.HandleFunc("/contacts", CreateContact(dbConn)).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", GetContact(dbConn)).Methods("GET")
	router.HandleFunc("/contacts/{id:[0-9]+}", UpdateContact(dbConn)).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", DeleteContact(dbConn)).Methods("DELETE")
}
