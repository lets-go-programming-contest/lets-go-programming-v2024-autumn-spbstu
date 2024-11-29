package controllers

import (
	"github.com/gorilla/mux"
)

func CreateRoutes(router *mux.Router) {
	router.HandleFunc("/contacts", GetContacts).Methods("GET")
}