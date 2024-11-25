package controllers

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts", CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", DeleteContact).Methods("DELETE")
	router.HandleFunc("/contacts/{id}/tags", GetTagsForContact).Methods("GET")
	router.HandleFunc("/contacts/{id}/tags", AddTagsToContact).Methods("POST")
	router.HandleFunc("/contacts/{id}/tags/{tag_id}", RemoveTagFromContact).Methods("DELETE")
	router.HandleFunc("/tags", GetTags).Methods("GET")
	router.HandleFunc("/contacts/{tag_id}", GetContactsByTag).Methods("GET")
}
