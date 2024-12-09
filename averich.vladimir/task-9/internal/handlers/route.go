package handlers

import (
	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router) {

	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", GetContact).Methods("GET")
	router.HandleFunc("/contacts", CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")

}
