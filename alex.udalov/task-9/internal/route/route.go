package route

import (
	"task-9/internal/contacts"

	"github.com/gorilla/mux"
)

func Create(router *mux.Router) {
	router.HandleFunc("/contacts", contacts.GetAll).Methods("GET")
	router.HandleFunc("/contacts/{id:[0-9]+}", contacts.Get).Methods("GET")
	router.HandleFunc("/contacts", contacts.Create).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", contacts.Update).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", contacts.Delete).Methods("DELETE")
}
