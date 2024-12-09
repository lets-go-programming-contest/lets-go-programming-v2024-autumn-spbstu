package manager

import (
	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router) {

	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id:[0-9]+}", GetContact).Methods("GET")
	router.HandleFunc("/contacts", CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", DeleteContact).Methods("DELETE")
}
