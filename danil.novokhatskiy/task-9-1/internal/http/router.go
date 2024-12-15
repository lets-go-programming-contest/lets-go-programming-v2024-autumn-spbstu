package http

import "github.com/gorilla/mux"

func CreateRoutes(r *mux.Router) {
	r.HandleFunc("/contacts/{id:[0-9]+}", GetContact).Methods("GET")
	r.HandleFunc("/contacts/{id:[0-9]+}", UpdateContact).Methods("PUT")
	r.HandleFunc("/contacts/{id:[0-9]+}", DeleteContact).Methods("DELETE")
	r.HandleFunc("/contacts", CreateContact).Methods("POST")
	r.HandleFunc("/contacts", GetContacts).Methods("GET")
	r.HandleFunc("/contacts", DeleteContacts).Methods("DELETE")
}
