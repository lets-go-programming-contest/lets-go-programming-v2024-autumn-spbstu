package route

import (
	"github.com/gorilla/mux"

	"github.com/Piyavva/task-9/internal/contact"
)

func Create(router *mux.Router) {
	router.HandleFunc("/contacts", contact.GetAll).Methods("GET")
	router.HandleFunc("/contacts/{id:[0-9]+}", contact.Get).Methods("GET")
	router.HandleFunc("/contacts", contact.Create).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", contact.Update).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", contact.Delete).Methods("DELETE")
}
