package route

import (
	"github.com/EmptyInsid/task-9/internal/handlers"
	"github.com/gorilla/mux"
)

func Setting(router *mux.Router) {
	router.HandleFunc("/contacts", handlers.GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", handlers.GetContact).Methods("GET")
	router.HandleFunc("/contacts", handlers.CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", handlers.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", handlers.DeleteContact).Methods("DELETE")
}
