package route

import (
	"github.com/KRYST4L614/task-9/internal/handler"
	"github.com/gorilla/mux"
)

func RegisterHandlers(router *mux.Router, handler handler.ContactHandler) {
	router.HandleFunc("/contacts/{id:[0-9]+}", handler.GetContact).Methods("GET")
	router.HandleFunc("/contacts", handler.GetAllContacts).Methods("GET")
	router.HandleFunc("/contacts", handler.AddContact).Methods("POST")
	router.HandleFunc("/contacts/{id:[0-9]+}", handler.UpdateContact).Methods("PUT")
	router.HandleFunc("/contacts/{id:[0-9]+}", handler.DeleteContactById).Methods("DELETE")
}
